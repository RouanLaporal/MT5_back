package web

import (
	"back_project/structure"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) AddShop() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		shop := structure.Shop{}
		json.NewDecoder(request.Body).Decode(&shop)

		id, err := h.Store.ShopStoreInterface.AddShop(shop)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(writer).Encode(struct {
			Status  string `json:"status"`
			Message string `json:"message"`
			NewShop int    `json:"newShop"`
		}{
			Status:  "success",
			Message: "Nouvelle boutique inséré avec succès",
			NewShop: id,
		})
	}
}

func (handler *Handler) GetAllShopByKindAndCity() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		QueryId := chi.URLParam(request, "id_kind")
		QueryCity := chi.URLParam(request, "city")
		id, _ := strconv.Atoi(QueryId)
		city := QueryCity
		writer.Header().Set("Content-Type", "application/json")
		shops, err := handler.Store.GetAllShopByKindAndCity(id, city)
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
