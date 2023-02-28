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

	handler.Post("/auth/login", handler.SignIn())
	handler.Post("/auth/register", handler.SignUp())

	/* Kind route */
	handler.Get("/get-kind", handler.GetKind())

	/* Shop route */

	handler.Post("/new-shop", handler.AddShop())
	handler.Get("/get-shop/{id_kind}/{city}", handler.GetAllShopByKindAndCity())
	handler.Get("/get-shop-by-user/{id_user}", handler.GetAllShopByUser())
	handler.Delete("/shop/{id}", handler.DeleteShop())
	handler.Patch("/shop/{id}", handler.UpdateShop())

	/* User route */
	handler.Post("/user/verify-password", middlewareCustom.IsAuthorized(handler.VerifyPassword()))
	handler.Patch("/user/update-password", middlewareCustom.IsAuthorized(handler.UpdatePassword()))
	handler.Patch("/user/update-profile", middlewareCustom.IsAuthorized(handler.UpdateUser()))
	handler.Delete("/user/delete-profile", middlewareCustom.IsAuthorized(handler.DeleteUser()))

	return handler
}

type Handler struct {
	*chi.Mux
	*database.Store
}
