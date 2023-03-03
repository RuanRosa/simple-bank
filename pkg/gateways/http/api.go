package http

import (
	"fmt"

	"github.com/RuanRosa/simple-bank/pkg/domain/usecase/account"
	"github.com/RuanRosa/simple-bank/pkg/domain/usecase/transfer"
	account_handler "github.com/RuanRosa/simple-bank/pkg/gateways/http/handler/account"
	auth_handler "github.com/RuanRosa/simple-bank/pkg/gateways/http/handler/auth"
	transfer_handler "github.com/RuanRosa/simple-bank/pkg/gateways/http/handler/transfer"

	auth_middleware "github.com/RuanRosa/simple-bank/pkg/gateways/http/middleware/auth"

	"github.com/RuanRosa/simple-bank/pkg/gateways/service/auth"

	"log"
	"net/http"

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

func (a *API) Start(
	accountUsecase account.IUsecase,
	authService auth.Service,
	authMiddleware auth_middleware.Middleware,
	transferUsecase transfer.IUsecase,
) {
	r := chi.NewRouter()

	r.Get("/health", a.healthCheck)

	account_handler.NewHandler(r, accountUsecase, authMiddleware)
	transfer_handler.NewHandler(r, transferUsecase, authMiddleware)
	auth_handler.NewHandler(r, authService)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", a.port), r); err != nil {
		log.Fatal(err)
	}
}

func (a *API) healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
