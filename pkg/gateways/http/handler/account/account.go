package account

import (
	"github.com/RuanRosa/simple-bank/pkg/domain/usecase/account"
	"github.com/go-chi/chi"
)

type Handler struct {
	usecase account.IUsecase
}

func NewHandler(c *chi.Mux, usecase account.IUsecase) Handler {
	h := Handler{
		usecase: usecase,
	}

	c.Route("/account", func(c chi.Router) {

		c.Get("/", h.All)
		c.Get("/{account_id}/balance", h.GetBalance)
		c.Post("/", h.Save)
	})

	return h
}
