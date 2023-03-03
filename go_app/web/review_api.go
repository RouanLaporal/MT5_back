package web

import (
	"back_project/helper"
	"back_project/structure"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// Review godoc
//
//	@Summary		AddReview
//	@Description	Create a review for a shop
//	@Tags			Review
//	@Accept			json
//	@Produce		json
//	@Param			id shop			body	integer		true	"shop id_shop"
//	@Param			id user			body	integer		true	"user id_user"
//	@Param			rating			body	string		false	"review rating"
//	@Param			comment			body	string		false	"review comment"
//	@Success		200
//	@Router			/review [post]
func (h *Handler) AddReview() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		token, err := helper.ExtractClaims(writer, request)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		review := structure.Review{}
		json.NewDecoder(request.Body).Decode(&review)

		id, err := h.Store.ReviewStoreInterface.AddReview(review, token.IDUser)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(writer).Encode(struct {
			Status    string `json:"status"`
			Message   string `json:"message"`
			NewReview int    `json:"newReview"`
		}{
			Status:    "success",
			Message:   "Nouveau commentaire ajouté avec succès",
			NewReview: id,
		})
	}
}

// Review godoc
//
//	@Summary		GetReviewByShop
//	@Description	Retrieve all reviews for a shop
//	@Tags			Review
//	@Accept			json
//	@Produce		json
//	@Param			id shop		path	integer		true	"shop id_shop"
//	@Success		200
//	@Router			/review/{id} [get]
func (h *Handler) GetReviewByShop() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		QueryId := chi.URLParam(request, "id")
		id, _ := strconv.Atoi(QueryId)
		writer.Header().Set("Content-Type", "application/json")
		reviews, err := h.Store.GetReviewByShop(id)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
		err = json.NewEncoder(writer).Encode(reviews)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// Review godoc
//
//	@Summary		UpdateReview
//	@Description	Update a review for a shop
//	@Tags			Review
//	@Accept			json
//	@Produce		json
//	@Param			id review		path	integer		true	"review id_review"
//	@Param			rating			body	string		false	"review rating"
//	@Param			comment			body	string		false	"review comment"
//	@Success		200
//	@Router			/review/{id} [patch]
func (h *Handler) UpdateReview() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		QueryId := chi.URLParam(request, "id")
		id, _ := strconv.Atoi(QueryId)
		review := structure.Review{}
		json.NewDecoder(request.Body).Decode(&review)
		writer.Header().Set("Content-Type", "application/json")
		err := h.Store.ReviewStoreInterface.UpdateReview(id, review)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(writer).Encode(struct {
			Status  string `json:"status"`
			Message string `json:"message"`
		}{
			Status:  "success",
			Message: "Commentaire modifié avec succès",
		})
	}
}

// Review godoc
//
//	@Summary		DeleteReview
//	@Description	Delete a review for a shop
//	@Tags			Review
//	@Accept			json
//	@Produce		json
//	@Param			id review	path	integer		true	"review id_review"
//	@Success		200
//	@Router			/review/{id} [delete]
func (h *Handler) DeleteReview() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		QueryId := chi.URLParam(request, "id")
		id, _ := strconv.Atoi(QueryId)
		writer.Header().Set("Content-Type", "application/json")
		err := h.Store.ReviewStoreInterface.DeleteReview(id)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(writer).Encode(true)
	}
}
