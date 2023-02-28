package web

import (
	"back_project/helper"
	"back_project/structure"
	"encoding/json"
	"net/http"
)

func (h *Handler) UpdateUser() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		token, err := helper.ExtractClaims(writer, request)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		user := structure.UpdateUser{}
		json.NewDecoder(request.Body).Decode(&user)

		error := h.Store.UserStoreInterface.UpdateUser(token.IDUser, user)

		if error != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(writer).Encode(true)
	}
}

func (h *Handler) VerifyPassword() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		password := structure.Password{}
		json.NewDecoder(request.Body).Decode(&password)
		claimsToken, err := helper.ExtractClaims(writer, request)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		auth, err := h.Store.UserStoreInterface.GetUserByEmail(claimsToken.Email)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		check := helper.CheckPasswordHash(password.Password, auth.Password)

		if !check {
			json.NewEncoder(writer).Encode(false)
			return
		}

		json.NewEncoder(writer).Encode(true)
	}
}

func (h *Handler) UpdatePassword() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		password := structure.Password{}
		json.NewDecoder(request.Body).Decode(&password)
		claimsToken, err := helper.ExtractClaims(writer, request)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		err = h.Store.UserStoreInterface.UpdatePassword(claimsToken.Email, password.Password)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(writer).Encode(true)
	}
}

func (h *Handler) DeleteUser() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		claimsToken, err := helper.ExtractClaims(writer, request)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		error := h.Store.UserStoreInterface.DeleteUser(claimsToken.IDUser)

		if error != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(writer).Encode(true)
	}
}
