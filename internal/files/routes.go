package files

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
)

type handler struct {
	db *sql.DB
}

func SetRoutes(r chi.Router, db *sql.DB) {

	h := handler{db}

	routes := func(r chi.Router) {
		r.Post("/", h.Modify)
	}

	r.Route("/files", routes)
}