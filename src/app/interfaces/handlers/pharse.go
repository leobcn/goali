package handlers

import (
	"app"
	"net/http"

	"github.com/alioygur/gores"
	"github.com/gorilla/mux"
)

var defaultLang = "en"

type pharseService interface {
	GetOrCreate(*app.Pharse) error
}

func NewPharse(srv pharseService, eh app.ErrorHandler) *Pharse {
	return &Pharse{srv, eh}
}

type Pharse struct {
	srv pharseService
	eh  app.ErrorHandler
}

func (p *Pharse) SetRoutes(r *mux.Router) {
	r.HandleFunc("/v1/{lang}/{pharse}", p.get).Methods("GET")
}

func (p *Pharse) get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	gc := func(map[string]string) (*app.Pharse, error) {
		pharse, err := app.NewPharse(params["pharse"])
		if err != nil {
			return nil, err
		}

		return pharse, p.srv.GetOrCreate(pharse)
	}

	if params["lang"] == defaultLang {
		gores.String(w, 200, params["pharse"])
		go gc(params)
		return
	}

	pharse, err := gc(params)
	if err != nil {
		p.eh.Handle(w, err)
		return
	}

	t, ok := pharse.Translations[params["lang"]]
	if !ok {
		gores.NoContent(w)
		return
	}

	gores.String(w, 200, t.Translate)
}
