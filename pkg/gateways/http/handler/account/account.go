package account

import (
    "github.com/go-chi/chi"
)

type Handler struct {
}

func NewHandler(c *chi.Mux) Handler {
    c.Route("/account", func(c chi.Router) {
        h := Handler{}

        c.Get("/", h.All)
        c.Get("/{account_id}/balance", h.GetBalance)
        c.Post("/", h.Save)
    })

    return Handler{}
}
