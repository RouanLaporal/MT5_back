package web

import (
	database "back_project/mysql"

	middlewareCustom "back_project/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func NewHandler(store *database.Store) *Handler {
	handler := &Handler{
		chi.NewRouter(),
		store,
	}

	handler.Use(middleware.Logger)

	handler.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	/* Authentification route */

	handler.Post("/login", handler.SignIn())
	handler.Post("/register", handler.SignUp())

	handler.Get("/get-kind", handler.GetKind())

	/* Shop route */

	handler.Post("/new-shop", handler.AddShop())
	handler.Get("/get-shop/{id_kind}/{city}", handler.GetAllShopByKindAndCity())
	handler.Get("/get-shop-by-user/{id_user}", handler.GetAllShopByUser())
	handler.Delete("/shop/{id}", handler.DeleteShop())
	handler.Patch("/shop/{id}", handler.UpdateShop())

	handler.Get("/collaborators/{id_shop}", handler.GetCollaboratorByShop())
	handler.Post("/new-collaborator", handler.AddCollaborator())
	handler.Patch("/collaborator/{id}", handler.UpdateCollaborator())
	handler.Delete("/collaborator/{id}", handler.DeleteCollaborator())

	handler.Post("/opening-hours", handler.AddOpeningHours())
	handler.Get("/opening-hours/{id_shop}", handler.GetOpeningHoursByShop())
	handler.Patch("/opening-hours/{id}", handler.UpdateOpeningHours())
	/* User route */
	handler.Delete("/delete/{id}", middlewareCustom.IsAuthorized(handler.DeleteUser()))
	handler.Patch("/user/update/{id}", handler.UpdateUser()) //modify for authorization

	return handler
}

type Handler struct {
	*chi.Mux
	*database.Store
}
