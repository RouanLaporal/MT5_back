package web

import (
	"back_project/helper"
	"back_project/structure"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// Shop godoc
//
//	@Summary		AddShop
//	@Description	create a shop and its informations
//	@Tags			Shop
//	@Accept			json
//	@Produce		json
//	@Param			name			body	string		false	"shop name"
//	@Param			zip_code		body	string		false	"shop zip_code"
//	@Param			city			body	string		false	"shop city"
//	@Param			lat				body	string		false	"shop lat"
//	@Param			long			body	string		false	"shop long"
//	@Param			country			body	string		false	"shop country"
//	@Param			phone			body	string		false	"shop phone"
//	@Param			email			body	string		false	"shop email"
//	@Param			description		body	string		false	"shop description"
//	@Param			kind id			body	integer		true	"kind id_kind"
//	@Param			id_user			body	integer		true	"user id_user"
//	@Success		200
//	@Router			/new-shop [post]
func (h *Handler) AddShopAndUser() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		shop := structure.NewShopAndUser{}
		json.NewDecoder(request.Body).Decode(&shop)
		_, err := h.Store.UserStoreInterface.GetUserByEmail(shop.UserEmail)
		if err == nil {
			http.Error(writer, "Email already in use", http.StatusBadRequest)
			return
		}
		err = h.Store.ShopStoreInterface.AddShopAndUser(shop)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		validToken, err := helper.GenerateJWT(shop.UserID, shop.UserEmail, "trader")
		if err != nil {
			http.Error(writer, "Failed to generate token", http.StatusInternalServerError)
			writer.Header().Set("Content-Type", "application/json")
			json.NewEncoder(writer).Encode(err)
			return
		}

		var authenticationUser structure.AuthUser

		authenticationUser.FirstName = shop.FirstName
		authenticationUser.LastName = shop.LastName
		authenticationUser.Email = shop.UserEmail
		authenticationUser.Phone = shop.UserPhone
		authenticationUser.Role = "trader"
		authenticationUser.TokenString = validToken

		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(authenticationUser)
	}
}

// func (h *Handler) AddShop() http.HandlerFunc {
// 	return func(writer http.ResponseWriter, request *http.Request) {
// 		token, err := helper.ExtractClaims(writer, request)
// 		if err != nil {
// 			http.Error(writer, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		shop := structure.NewShop{}
// 		json.NewDecoder(request.Body).Decode(&shop)

// 		id, err := h.Store.ShopStoreInterface.AddShop(shop, token.IDUser)

// 		if err != nil {
// 			http.Error(writer, err.Error(), http.StatusInternalServerError)
// 			return
// 		}

// 		json.NewEncoder(writer).Encode(struct {
// 			Status  string `json:"status"`
// 			Message string `json:"message"`
// 			NewShop int    `json:"newShop"`
// 		}{
// 			Status:  "success",
// 			Message: "Nouveau commentaire ajouté avec succès",
// 			NewShop: id,
// 		})
// 	}
// }

// Shop godoc
//
//	@Summary		GetAllShopByKindAndCity
//	@Description	Retrieve all shop by kind and city
//	@Tags			Shop
//	@Accept			json
//	@Produce		json
//	@Param			city			path	string		true	"shop city"
//	@Param			kind id			path	integer		true	"kind id_kind"
//	@Success		200
//	@Router			/get-shop/{id_kind}/{city} [get]
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

// Shop godoc
//
//	@Summary		DeleteShop
//	@Description	Delete a shop by its id
//	@Tags			Shop
//	@Produce		json
//	@Param			id_shop		path	integer		true	"shop id_shop"
//	@Success		200
//	@Router			/shop/{id_shop} [delete]
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

// Shop godoc
//
//	@Summary		UpdateShop
//	@Description	Update a shop by its id
//	@Tags			Shop
//	@Accept			json
//	@Produce		json
//	@Param			name			body	string		false	"shop name"
//	@Param			zip_code		body	string		false	"shop zip_code"
//	@Param			city			body	string		false	"shop city"
//	@Param			lat				body	string		false	"shop lat"
//	@Param			long			body	string		false	"shop long"
//	@Param			country			body	string		false	"shop country"
//	@Param			phone			body	string		false	"shop phone"
//	@Param			email			body	string		false	"shop email"
//	@Param			description		body	string		false	"shop description"
//	@Success		200
//	@Router			/shop/{id} [patch]
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

// Shop godoc
//
//	@Summary		GetAllShopByUser
//	@Description	Retrieve all shop by user
//	@Tags			Shop
//	@Produce		json
//	@Header			200		{string} 	Token	"user token"
//	@Param			user id		path	integer		true	"user id_user"
//	@Success		200
//	@Router			/get-shop-by-user [get]
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

// Shop godoc
//
//	@Summary		GetAllShopNear
//	@Description	Retrieve all shop near the area defined by latitude and longitude
//	@Tags			Shop
//	@Produce		json
//	@Param			lng			path	integer		true	"shop long"
//	@Param			lat			path	integer		true	"shop lat"
//	@Param			kind id		path	integer		true	"kind id_kind"
//	@Success		200
//	@Router			/shops/nearby/{lng}/{lat}/{kind} [get]
func (handler *Handler) GetAllShopNear() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		QueryLng := chi.URLParam(request, "lng")
		QueryLat := chi.URLParam(request, "lat")
		QueryKind := chi.URLParam(request, "kind")
		lng, _ := strconv.ParseFloat(QueryLng, 64)
		lat, _ := strconv.ParseFloat(QueryLat, 64)
		kind := QueryKind
		fmt.Println(lng)
		fmt.Println(lat)
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

// Shop godoc
//
//	@Summary		GetShop
//	@Description	Retrieve all shop by id shope
//	@Tags			Shop
//	@Accept			json
//	@Produce		json
//	@Param			id shop		path	integer		true	"shop id_shop"
//	@Success		200
//	@Router			/shop/{id} [get]
func (handler *Handler) GetShop() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		QueryId := chi.URLParam(request, "id")
		id, _ := strconv.Atoi(QueryId)
		writer.Header().Set("Content-Type", "application/json")
		shop, err := handler.Store.GetShopById(id)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
		err = json.NewEncoder(writer).Encode(shop)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
