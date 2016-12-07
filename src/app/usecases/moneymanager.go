package usecases

import (
	"app"
	"app/interfaces/errs"
	"time"
)

var (
	errInvalidUser = errs.BadRequest("invalid user")
)

type mmRepo interface {
	app.IDatabase
	GetTransactions(app.DBWhere, *app.DBFilter) ([]app.Transaction, error)
	GetAccounts(app.DBWhere, *app.DBFilter) ([]app.Account, error)
}

func NewMoneyManager(r mmRepo) *MoneyManager {
	return &MoneyManager{r}
}

type MoneyManager struct {
	mmRepo
}

func (mm *MoneyManager) AddTransaction(f *TransactionForm) (*app.Transaction, error) {
	// check UserID
	exists, err := mm.ExistsBy(&app.User{}, app.DBWhere{"id": f.UserID})
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, errs.Wrap(errInvalidUser)
	}

	// todo: check AccountID

	var tm app.Transaction
	tm.UserID = f.UserID
	tm.AccountID = f.AccountID
	tm.Amount = f.Amount
	tm.Desc = f.Desc

	if f.Date == nil {
		tm.Date = time.Now()
	} else {
		tm.Date = *f.Date
	}

	tm.AddTag(f.Tags...)

	for i, c := range tm.Tags {
		if err := mm.FirstOrInit(&c, app.DBWhere{"name": c.Name}); err != nil {
			return nil, err
		}
		tm.Tags[i] = c
	}

	return &tm, mm.Store(&tm)
}

func (mm *MoneyManager) AddAccount(f *AccountForm) (*app.Account, error) {
	var a app.Account
	a.Name = f.Name
	a.Desc = f.Desc
	a.Type = f.Type

	if err := mm.Store(&a); err != nil {
		return nil, err
	}

	return &a, nil
}

type TransactionForm struct {
	UserID    int
	AccountID int
	Amount    float32
	Desc      string
	Tags      []string
	Date      *time.Time
}

type AccountForm struct {
	Owner int
	Name  string
	Desc  string
	Type  string
}
