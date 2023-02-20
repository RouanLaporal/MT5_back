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

	/* Authentification route */

	handler.Post("/login", handler.SignIn())
	handler.Post("/register", handler.SignUp())

	handler.Get("/get-kind", handler.GetKind())

	/* Shop toute */

	handler.Post("/new-shop", handler.AddShop())
	handler.Get("/get-shop/{id_kind}/{city}", handler.GetAllShopByKindAndCity())
	handler.Delete("/shop/{id}", handler.DeleteShop())
	handler.Patch("/shop/{id}", handler.UpdateShop())

	// handler.Delete("/delete/{id}", middlewareCustom.IsAuthorized(handler.DeleteUser()))

	return handler
}

type Handler struct {
	*chi.Mux
	*database.Store
}
