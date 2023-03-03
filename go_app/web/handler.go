package web

import (
	database "back_project/mysql"

	middlewareCustom "back_project/middleware"

	_ "back_project/docs"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
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

	handler.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8097/swagger/doc.json"), //The url pointing to API definition
	))
	/* Authentification route */

	handler.Post("/auth/login", handler.SignIn())
	handler.Post("/auth/register", handler.SignUp())

	/* Shop route */

	handler.Post("/new-shop", handler.AddShopAndUser())
	handler.Get("/get-shop-by-user", middlewareCustom.IsAuthorized(handler.GetAllShopByUser()))
	handler.Get("/shop/{id}", handler.GetShop())
	handler.Delete("/shop/{id}", middlewareCustom.IsAuthorized(handler.DeleteShop()))
	handler.Patch("/shop/{id}", middlewareCustom.IsAuthorized(handler.UpdateShop()))

	handler.Get("/shops/nearby/{lng}/{lat}/{kind}", handler.GetAllShopNear())
	handler.Get("/get-shop/{id_kind}/{city}", handler.GetAllShopByKindAndCity())
	handler.Get("/kinds", handler.GetKind())

	handler.Get("/collaborators/{id_shop}", handler.GetCollaboratorByShop())
	handler.Post("/new-collaborator", handler.AddCollaborator())
	handler.Patch("/collaborator/{id}", handler.UpdateCollaborator())
	handler.Delete("/collaborator/{id}", handler.DeleteCollaborator())

	handler.Post("/opening-hours", middlewareCustom.IsAuthorized(handler.AddOpeningHours()))
	handler.Patch("/opening-hours/{id}", middlewareCustom.IsAuthorized(handler.UpdateOpeningHours()))
	handler.Delete("/opening-hours/{id}", middlewareCustom.IsAuthorized(handler.DeleteOpeningHours()))
	handler.Get("/opening-hours/{id_shop}", handler.GetOpeningHoursByShop())

	/* User route */
	handler.Post("/user/verify-password", middlewareCustom.IsAuthorized(handler.VerifyPassword()))
	handler.Patch("/user/update-password", middlewareCustom.IsAuthorized(handler.UpdatePassword()))
	handler.Patch("/user/update-profile", middlewareCustom.IsAuthorized(handler.UpdateUser()))
	handler.Delete("/user/delete-profile", middlewareCustom.IsAuthorized(handler.DeleteUser()))

	/* Benefit routes */
	handler.Get("/benefit/{id}", handler.GetBenefitByShop())
	handler.Post("/benefit", middlewareCustom.IsAuthorized(handler.AddBenefit()))
	handler.Patch("/benefit/{id}", middlewareCustom.IsAuthorized(handler.UpdateBenefit()))
	handler.Delete("/benefit/{id}", middlewareCustom.IsAuthorized(handler.DeleteBenefit()))

	/* Review routes */
	handler.Get("/review/{id}", handler.GetReviewByShop())
	handler.Post("/review", middlewareCustom.IsAuthorized(handler.AddReview()))
	handler.Patch("/review/{id}", middlewareCustom.IsAuthorized(handler.UpdateReview()))
	handler.Delete("/review/{id}", middlewareCustom.IsAuthorized(handler.DeleteReview()))

	/*Reservation routes*/
	handler.Post("/reservation", middlewareCustom.IsAuthorized(handler.AddReservation()))
	handler.Get("/reservation/{id_shop}", (handler.GetExistingReservationForPeriod()))
	handler.Get("/reservation", middlewareCustom.IsAuthorized(handler.GetReservationByUser()))
	handler.Patch("/reservation/{id}", middlewareCustom.IsAuthorized(handler.UpdateReservation()))
	handler.Delete("/reservation/{id}", middlewareCustom.IsAuthorized(handler.DeleteReservation()))

	return handler
}

type Handler struct {
	*chi.Mux
	*database.Store
}
