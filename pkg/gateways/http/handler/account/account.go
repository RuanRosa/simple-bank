package account

import (
	"github.com/RuanRosa/simple-bank/pkg/domain/usecase/account"
	"github.com/RuanRosa/simple-bank/pkg/gateways/http/middleware/auth"
	"github.com/go-chi/chi"
)

type Handler struct {
	usecase        account.IUsecase
	authMiddleware auth.Middleware
}

func NewHandler(c *chi.Mux, usecase account.IUsecase, authMiddleware auth.Middleware) Handler {
	h := Handler{
		usecase:        usecase,
		authMiddleware: authMiddleware,
	}

	c.Group(func(c chi.Router) {
		c.Use(authMiddleware.Check)
		c.Route("/account", func(c chi.Router) {
			c.Get("/", h.All)
			c.Get("/{account_id}/balance", h.GetBalance)
			c.Post("/", h.Save)
		})
	})

	return h
}
