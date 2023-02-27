package web

import (
	"back_project/structure"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (handler *Handler) AddOpeningHours() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		opening_hours := structure.OpeningHours{}
		json.NewDecoder(request.Body).Decode(&opening_hours)

		id, err := handler.Store.OpeningHoursStoreInterface.AddOpeningHours(opening_hours)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(writer).Encode(struct {
			Status          string `json:"status"`
			Message         string `json:"message"`
			NewOpeningHours int    `json:"newOpeningHours"`
		}{
			Status:          "success",
			Message:         "Heure d'ouverture bien enregistré",
			NewOpeningHours: id,
		})
	}
}

func (handler *Handler) GetOpeningHoursByShop() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		QueryId := chi.URLParam(request, "id_shop")
		id_shop, _ := strconv.Atoi(QueryId)
		writer.Header().Set("Content-Type", "application/json")
		opening_hours, err := handler.Store.GetOpeningHoursByShop(id_shop)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
		err = json.NewEncoder(writer).Encode(opening_hours)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
