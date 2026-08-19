package users

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
)

type handler struct {
	db *sql.DB
}

func SetRoutes(r chi.Router, db *sql.DB) {
	h := handler{db}

	// Handler.HTTPVerb have same method signature as Handler.ServeHTTP
	r.Post("/", h.Create)
	r.Delete("/{id}", h.Delete)
}
