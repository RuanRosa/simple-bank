package transfer

import (
    "github.com/go-chi/chi"
)

type Handler struct {
}

func NewHandler(c *chi.Mux) Handler {
    c.Route("/transfer", func(c chi.Router) {
        h := Handler{}

        c.Get("/", h.All)
        c.Post("/", h.Save)
    })

    return Handler{}
}
