package web

import (
	database "back_project/mysql"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewHandler(store *database.Store) *Handler {
	handler := &Handler{
		chi.NewRouter(),
		store,
	}

	handler.Use(middleware.Logger)

	handler.Post("/login", handler.SignIn())
	handler.Post("/register", handler.SignUp())

	return handler
}

type Handler struct {
	*chi.Mux
	*database.Store
}
