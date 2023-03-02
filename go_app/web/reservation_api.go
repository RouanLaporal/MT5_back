package web

import (
	"back_project/helper"
	"back_project/structure"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (handler *Handler) AddReservation() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		token, err := helper.ExtractClaims(writer, request)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		reservation := structure.Reservation{}
		json.NewDecoder(request.Body).Decode(&reservation)

		id, err := handler.Store.ReservationStoreInterface.AddReservation(reservation, token.IDUser)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(writer).Encode(struct {
			Status         string `json:"status"`
			Message        string `json:"message"`
			NewReservation int    `json:"newReservation"`
		}{
			Status:         "success",
			Message:        "Nouvelle réservation ajouté avec succès",
			NewReservation: id,
		})
	}
}

func (handler *Handler) GetExistingReservationForPeriod() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		QueryId := chi.URLParam(request, "id_shop")
		id, _ := strconv.Atoi(QueryId)
		writer.Header().Set("Content-Type", "application/json")
		shops, err := handler.Store.GetExistingReservationForPeriod(id)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
		err = json.NewEncoder(writer).Encode(shops)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (handler *Handler) GetReservationByUser() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		token, err := helper.ExtractClaims(writer, request)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		writer.Header().Set("Content-Type", "application/json")
		reservations, err := handler.Store.GetReservationByUser(token.IDUser)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
		err = json.NewEncoder(writer).Encode(reservations)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (handler *Handler) UpdateReservation() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		QueryId := chi.URLParam(request, "id")
		id, _ := strconv.Atoi(QueryId)
		reservation := structure.UpdateReservation{}
		json.NewDecoder(request.Body).Decode(&reservation)
		writer.Header().Set("Content-Type", "application/json")
		err := handler.Store.ReservationStoreInterface.UpdateReservation(id, reservation)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(writer).Encode(true)

	}
}

func (handler *Handler) DeleteReservation() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		QueryId := chi.URLParam(request, "id")
		id, _ := strconv.Atoi(QueryId)
		writer.Header().Set("Content-Type", "application/json")
		err := handler.Store.ReservationStoreInterface.DeleteReservation(id)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(writer).Encode(true)
	}
}
