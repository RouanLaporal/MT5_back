package web

import (
	"back_project/structure"
	"encoding/json"
	"net/http"
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
			NewOpeningHours int    `json:"newShop"`
		}{
			Status:          "success",
			Message:         "Heure d'ouverture bien enregistré",
			NewOpeningHours: id,
		})
	}
}
