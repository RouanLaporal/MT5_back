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

		validToken, err := helper.GenerateJWT(auth.Email, auth.Role)
		if err != nil {
			http.Error(writer, "Failed to generate token", http.StatusInternalServerError)
			writer.Header().Set("Content-Type", "application/json")
			json.NewEncoder(writer).Encode(err)
			return
		}

		var token structure.Token
		token.Email = auth.Email
		token.Role = auth.Role
		token.TokenString = validToken
		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(token)
	}
}

func (h *Handler) SignUp() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		user := structure.User{}
		json.NewDecoder(request.Body).Decode(&user)

		_, err := h.Store.UserStoreInterface.GetUserByEmail(user.Email)
		if err == nil {
			http.Error(writer, "Email already in use", http.StatusBadRequest)
			return
		}

		id, err := h.Store.UserStoreInterface.AddUser(user)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(writer).Encode(struct {
			Status  string `json:"status"`
			Message string `json:"message"`
			NewUser int    `json:"newUser"`
		}{
			Status:  "success",
			Message: "Nouveau user inséré avec succès",
			NewUser: id,
		})
	}
}
