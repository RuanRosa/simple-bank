package http

import (
    "github.com/RuanRosa/simple-bank/pkg/gateways/http/handler/account"
    "log"
	"net/http"

	"github.com/go-chi/chi"
)

type API struct{}

func NewAPI() *API {
	return &API{}
}

func (a *API) Start() {
	r := chi.NewRouter()

    account.NewHandler(r)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
