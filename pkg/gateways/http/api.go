package http

import (
	"fmt"

	"github.com/RuanRosa/simple-bank/pkg/domain/usecase/account"
	account_handler "github.com/RuanRosa/simple-bank/pkg/gateways/http/handler/account"

	"log"
	"net/http"

	"github.com/RuanRosa/simple-bank/pkg/gateways/http/handler/transfer"

	"github.com/go-chi/chi"
)

type API struct {
	port string
}

func NewAPI(port string) *API {
	return &API{
		port: port,
	}
}

func (a *API) Start(usecase account.IUsecase) {
	r := chi.NewRouter()
	r.Get("/health", a.healthCheck)

	account_handler.NewHandler(r, usecase)
	transfer.NewHandler(r)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", a.port), r); err != nil {
		log.Fatal(err)
	}
}

func (a *API) healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
