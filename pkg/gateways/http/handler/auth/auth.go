package auth

import (
	"github.com/RuanRosa/simple-bank/pkg/gateways/service/auth"
	"github.com/go-chi/chi"
)

type Handler struct {
	service auth.Service
}

func NewHandler(c *chi.Mux, service auth.Service) Handler {
	h := Handler{
		service: service,
	}

	c.Post("/login", h.Login)

	return h
}
