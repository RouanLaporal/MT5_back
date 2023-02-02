package web

import (
	"back_project/helper"
	"back_project/structure"
	"encoding/json"
	"net/http"
)

func (h *Handler) SignIn() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		authentication := structure.Authentication{}
		json.NewDecoder(request.Body).Decode(&authentication)
		auth, err := h.Store.UserStoreInterface.GetUserByEmail(authentication.Email)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		check := helper.CheckPasswordHash(authentication.Password, auth.Password)

		if !check {
			http.Error(writer, "Username or Password is incorrect", http.StatusBadRequest)
			writer.Header().Set("Content-Type", "application/json")
			json.NewEncoder(writer).Encode(err)
			return
		}
	}
}
