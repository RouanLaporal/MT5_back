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
		AllowCredentials: true,
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"X-PINGOTHER", "Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Token", "Access-Control-Allow-Origin"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
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
	/* Benefit routes */
	handler.Get("/benefit/get/{id}", handler.GetBenefitByShop())
	handler.Post("/benefit/add", middlewareCustom.IsAuthorized(handler.AddBenefit()))
	handler.Patch("/benefit/update/{id}", middlewareCustom.IsAuthorized(handler.UpdateBenefit()))
	handler.Delete("/benefit/delete/{id}", middlewareCustom.IsAuthorized(handler.DeleteBenefit()))

	/* Review routes */
	handler.Get("/review/get/{id}", handler.GetReviewByShop())
	handler.Post("/review/add", handler.AddReview())
	handler.Patch("/review/update/{id}", handler.UpdateReview())
	handler.Delete("/review/delete/{id}", handler.DeleteReview())

	return handler
}

type Handler struct {
	*chi.Mux
	*database.Store
}
