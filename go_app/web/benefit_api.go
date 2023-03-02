package web

import (
	"back_project/structure"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) AddBenefit() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		benefit := structure.Benefit{}
		json.NewDecoder(request.Body).Decode(&benefit)

		id, err := h.Store.BenefitStoreInterface.AddBenefit(benefit)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(writer).Encode(struct {
			Status     string `json:"status"`
			Message    string `json:"message"`
			NewBenefit int    `json:"newBenefit"`
		}{
			Status:     "success",
			Message:    "Nouvelle prestation insérée avec succès",
			NewBenefit: id,
		})
	}
}

func (h *Handler) GetBenefitByShop() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		QueryId := chi.URLParam(request, "id")
		id, _ := strconv.Atoi(QueryId)
		writer.Header().Set("Content-Type", "application/json")
		benefits, err := h.Store.GetBenefitByShop(id)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
		err = json.NewEncoder(writer).Encode(benefits)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (h *Handler) UpdateBenefit() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		QueryId := chi.URLParam(request, "id")
		id, _ := strconv.Atoi(QueryId)
		benefit := structure.Benefit{}
		json.NewDecoder(request.Body).Decode(&benefit)
		writer.Header().Set("Content-Type", "application/json")
		err := h.Store.BenefitStoreInterface.UpdateBenefit(id, benefit)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(writer).Encode(struct {
			Status  string `json:"status"`
			Message string `json:"message"`
		}{
			Status:  "success",
			Message: "Prestation modifiée avec succès",
		})
	}
}

func (h *Handler) DeleteBenefit() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		QueryId := chi.URLParam(request, "id")
		id, _ := strconv.Atoi(QueryId)
		writer.Header().Set("Content-Type", "application/json")
		err := h.Store.BenefitStoreInterface.DeleteBenefit(id)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(writer).Encode(struct {
			Status  string `json:"status"`
			Message string `json:"message"`
		}{
			Status:  "success",
			Message: "Prestation supprimée avec succès",
		})
	}
}
