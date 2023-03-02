package web

import (
	"back_project/structure"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// Opening hours godoc
//
//	@Summary		AddOpeningHours
//	@Description	Create a range of opening hours for a shop
//	@Tags			Opening hours
//	@Accept			json
//	@Produce		json
//	@Param			id day		body	integer		true	"opening id_day"
//	@Param			id shop		body	integer		true	"opening id_shop"
//	@Param			open		body	string		true	"opening open"
//	@Param			close		body	string		true	"opening close"
//	@Success		200
//	@Router			/opening-hours [post]
func (handler *Handler) AddOpeningHours() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		opening_hours := structure.OpeningHours{}
		json.NewDecoder(request.Body).Decode(&opening_hours)

		_, err := handler.Store.OpeningHoursStoreInterface.AddOpeningHours(opening_hours)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(writer).Encode(true)
	}
}

// Opening hours godoc
//
//	@Summary		GetOpeningHoursByShop
//	@Description	Retrieve opening hours for a shop
//	@Tags			Opening hours
//	@Accept			json
//	@Produce		json
//	@Param			id shop		path	integer		true	"shop id_shop"
//	@Success		200
//	@Router			/opening-hours/{id_shop} [get]
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

// Opening hours godoc
//
//	@Summary		UpdateOpeningHours
//	@Description	Update opening hours for a shop
//	@Tags			Opening hours
//	@Accept			json
//	@Produce		json
//	@Param			id opening hour		path	integer		true	"opening hour id"
//	@Success		200
//	@Router			/opening-hours/{id} [patch]
func (handler *Handler) UpdateOpeningHours() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		QueryId := chi.URLParam(request, "id")
		id, _ := strconv.Atoi(QueryId)
		opening_hours := structure.OpeningHours{}

		json.NewDecoder(request.Body).Decode(&opening_hours)
		writer.Header().Set("Content-Type", "application/json")
		err := handler.Store.OpeningHoursStoreInterface.UpdateOpeningHours(id, opening_hours)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(writer).Encode(true)
	}
}

// Opening hours godoc
//
//	@Summary		DeleteOpeningHours
//	@Description	Delete an opening hours for a shop
//	@Tags			Opening hours
//	@Accept			json
//	@Produce		json
//	@Param			id opening hour		path	integer		true	"opening hour id"
//	@Success		200
//	@Router			/opening-hours/{id} [delete]
func (handler *Handler) DeleteOpeningHours() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		QueryId := chi.URLParam(request, "id")
		id, _ := strconv.Atoi(QueryId)
		writer.Header().Set("Content-Type", "application/json")
		err := handler.Store.OpeningHoursStoreInterface.DeleteOpeningHours(id)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(writer).Encode(true)
	}
}
