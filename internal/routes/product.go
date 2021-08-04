package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Product struct {
	Get http.HandlerFunc
}

func (g *Product) Install(r *chi.Mux) {
	r.Get("/products/{product_id}", g.Get)
}
