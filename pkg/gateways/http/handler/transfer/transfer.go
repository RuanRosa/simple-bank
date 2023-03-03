package transfer

import (
	"github.com/RuanRosa/simple-bank/pkg/domain/usecase/transfer"
	"github.com/RuanRosa/simple-bank/pkg/gateways/http/middleware/auth"
	"github.com/go-chi/chi"
)

type Handler struct {
	usecase        transfer.IUsecase
	authMiddleware auth.Middleware
}

func NewHandler(
	c *chi.Mux,
	usecase transfer.IUsecase,
	authMiddleware auth.Middleware,
) Handler {
	h := Handler{
		usecase:        usecase,
		authMiddleware: authMiddleware,
	}

	c.Group(func(c chi.Router) {
		c.Use(authMiddleware.Check)
		c.Route("/transfer", func(c chi.Router) {

			c.Get("/", h.All)
			c.Post("/", h.TransferMoney)
		})
	})

	return Handler{}
}
