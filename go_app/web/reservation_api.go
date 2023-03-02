package web

import (
	"back_project/helper"
	"back_project/structure"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// Reservation godoc
//
//	@Summary		AddReservation
//	@Description	Create a reservation for a shop
//	@Tags			Reservation
//	@Accept			json
//	@Produce		json
//	@Param			id shop			body	integer		true	"shop id_shop"
//	@Param			id user			body	integer		true	"user id_user"
//	@Param			id benefit		body	integer		true	"benefit id_benefit"
//	@Param			date			body	string		true	"reservation date"
//	@Param			time			body	string		true	"reservation time"
//	@Param			comment			body	string		false	"reservation comment"
//	@Success		200
//	@Router			/reservation [post]
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

// Reservation godoc
//
//	@Summary		GetExistingReservationForPeriod
//	@Description	Retrieve all existing reservations for a period in the last 90 days
//	@Tags			Reservation
//	@Accept			json
//	@Produce		json
//	@Param			id shop			path	integer		true	"shop id_shop"
//	@Success		200
//	@Router			/reservation/{id_shop} [get]
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

// Reservation godoc
//
//	@Summary		GetReservationByUser
//	@Description	Retrieve all existing reservations for a user
//	@Tags			Reservation
//	@Produce		json
//	@Param			id user		body	integer		true	"user id_user"
//	@Success		200
//	@Router			/reservation [get]
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

// Reservation godoc
//
//	@Summary		UpdateReservation
//	@Description	Update a reservation in a shop
//	@Tags			Reservation
//	@Accept			json
//	@Produce		json
//	@Param			id reservation			path	integer		true	"user id_reservation"
//	@Param			id benefit		body	integer		false	"benefit id_benefit"
//	@Param			date			body	string		false	"reservation date"
//	@Param			time			body	string		false	"reservation time"
//	@Param			comment			body	string		false	"reservation comment"
//	@Success		200
//	@Router			/reservation/{id} [patch]
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

// Reservation godoc
//
//	@Summary		DeleteReservation
//	@Description	Delete a reservation in a shop
//	@Tags			Reservation
//	@Accept			json
//	@Produce		json
//	@Param			id reservation		path	integer		true	"user id_reservation"
//	@Success		200
//	@Router			/reservation/{id} [delete]
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
