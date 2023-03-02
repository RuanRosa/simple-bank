package http

import (
	"fmt"

	"github.com/RuanRosa/simple-bank/pkg/domain/usecase/account"
	account_handler "github.com/RuanRosa/simple-bank/pkg/gateways/http/handler/account"
	auth_handler "github.com/RuanRosa/simple-bank/pkg/gateways/http/handler/auth"
	"github.com/RuanRosa/simple-bank/pkg/gateways/service/auth"

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

func (a *API) Start(usecase account.IUsecase, authService auth.Service) {
	r := chi.NewRouter()
	r.Get("/health", a.healthCheck)

	account_handler.NewHandler(r, usecase)
	transfer.NewHandler(r)
	auth_handler.NewHandler(r, authService)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", a.port), r); err != nil {
		log.Fatal(err)
	}
}

func (a *API) healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
