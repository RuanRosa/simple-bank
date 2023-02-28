package http

import (
    "github.com/RuanRosa/simple-bank/pkg/gateways/http/handler/account"
    "github.com/RuanRosa/simple-bank/pkg/gateways/http/handler/transfer"
    "log"
	"net/http"

	"github.com/go-chi/chi"
)

type API struct{
    port string
}

func NewAPI(port string) *API {
	return &API{
        port: port,
    }
}

func (a *API) Start() {
	r := chi.NewRouter()

    account.NewHandler(r)
    transfer.NewHandler(r)

	if err := http.ListenAndServe(a.port, r); err != nil {
		log.Fatal(err)
	}
}
