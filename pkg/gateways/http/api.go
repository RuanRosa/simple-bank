package http

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type API struct{}

func NewAPI() *API {
	return &API{}
}

func (a *API) Start() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
