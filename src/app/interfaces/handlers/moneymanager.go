package handlers

import (
	"app"
	"app/usecases"
	"net/http"

	"github.com/alioygur/gores"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

type transactionService interface {
	AddTransaction(*usecases.TransactionForm) (*app.Transaction, error)
	GetTransactions(app.DBWhere, *app.DBFilter) ([]app.Transaction, error)
	AddAccount(*usecases.AccountForm) (*app.Account, error)
	GetAccounts(app.DBWhere, *app.DBFilter) ([]app.Account, error)
}

func NewMoneyManager(srv transactionService, eh app.ErrorHandler) *Expense {
	return &Expense{srv, eh}
}

type Expense struct {
	srv transactionService
	eh  app.ErrorHandler
}

func (e *Expense) SetRoutes(r *mux.Router, mid ...alice.Constructor) {
	h := alice.New(mid...)
	r.Handle("/v1/transactions", h.ThenFunc(e.addTransaction)).Methods("POST")
	r.Handle("/v1/transactions", h.ThenFunc(e.getTransactions)).Methods("GET")
	r.Handle("/v1/accounts", h.ThenFunc(e.addAccount)).Methods("POST")
	r.Handle("/v1/accounts", h.ThenFunc(e.getAccounts)).Methods("GET")
}

func (e *Expense) addTransaction(w http.ResponseWriter, r *http.Request) {
	f := new(usecases.TransactionForm)
	if err := decodeR(r, f); err != nil {
		e.eh.Handle(w, err)
		return
	}

	usr := app.UserMustFromContext(r.Context())
	f.UserID = usr.ID

	em, err := e.srv.AddTransaction(f)
	if err != nil {
		e.eh.Handle(w, err)
		return
	}

	gores.JSON(w, 200, response{em})
}

func (e *Expense) getTransactions(w http.ResponseWriter, r *http.Request) {
	usr := app.UserMustFromContext(r.Context())
	fi := app.DBFilter{Preload: []string{"Tags", "Account"}}
	es, err := e.srv.GetTransactions(app.DBWhere{"user_id": usr.ID}, &fi)
	if err != nil {
		e.eh.Handle(w, err)
		return
	}

	gores.JSON(w, 200, response{es})
}

func (e *Expense) addAccount(w http.ResponseWriter, r *http.Request) {
	usr := app.UserMustFromContext(r.Context())

	f := new(usecases.AccountForm)
	if err := decodeR(r, f); err != nil {
		e.eh.Handle(w, err)
		return
	}
	f.Owner = usr.ID
	a, err := e.srv.AddAccount(f)
	if err != nil {
		e.eh.Handle(w, err)
		return
	}
	gores.JSON(w, 201, response{a})
}

func (e *Expense) getAccounts(w http.ResponseWriter, r *http.Request) {
	usr := app.UserMustFromContext(r.Context())
	as, err := e.srv.GetAccounts(app.DBWhere{"user_id": usr.ID}, nil)
	if err != nil {
		e.eh.Handle(w, err)
		return
	}

	gores.JSON(w, 200, response{as})
}
