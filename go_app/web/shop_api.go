package web

import (
	"back_project/helper"
	"back_project/structure"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) AddShop() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		token, err := helper.ExtractClaims(writer, request)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		shop := structure.NewShop{}
		json.NewDecoder(request.Body).Decode(&shop)

		id, err := h.Store.ShopStoreInterface.AddShop(shop, token.IDUser)

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
			Message: "Nouveau commentaire ajouté avec succès",
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

func (handler *Handler) DeleteShop() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		QueryId := chi.URLParam(request, "id")
		id, _ := strconv.Atoi(QueryId)
		writer.Header().Set("Content-Type", "application/json")
		err := handler.Store.ShopStoreInterface.DeleteShop(id)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(writer).Encode(true)
	}
}

func (handler *Handler) UpdateShop() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		QueryId := chi.URLParam(request, "id")
		id, _ := strconv.Atoi(QueryId)
		shop := structure.Shop{}
		json.NewDecoder(request.Body).Decode(&shop)
		writer.Header().Set("Content-Type", "application/json")
		err := handler.Store.ShopStoreInterface.UpdateShop(id, shop)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(writer).Encode(true)

	}
}

func (handler *Handler) GetAllShopByUser() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		token, err := helper.ExtractClaims(writer, request)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		writer.Header().Set("Content-Type", "application/json")
		shops, err := handler.Store.GetAllShopByUser(token.IDUser)
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

func (handler *Handler) GetAllShopNear() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		QueryLng := chi.URLParam(request, "lng")
		QueryLat := chi.URLParam(request, "lat")
		QueryKind := chi.URLParam(request, "kind")
		lng, _ := strconv.ParseFloat(QueryLng, 64)
		lat, _ := strconv.ParseFloat(QueryLat, 64)
		kind := QueryKind
		writer.Header().Set("Content-Type", "application/json")
		shops, err := handler.ShopStoreInterface.GetAllShopNear(lat, lng, kind)
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
