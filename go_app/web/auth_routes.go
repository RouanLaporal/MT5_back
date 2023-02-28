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

		validToken, err := helper.GenerateJWT(auth.ID, auth.Email, auth.Role)
		if err != nil {
			http.Error(writer, "Failed to generate token", http.StatusInternalServerError)
			writer.Header().Set("Content-Type", "application/json")
			json.NewEncoder(writer).Encode(err)
			return
		}

		var token structure.Token
		var authenticationUser structure.AuthUser

		token.Email = auth.Email
		token.Role = auth.Role
		token.TokenString = validToken
		authenticationUser.FirstName = auth.FirstName
		authenticationUser.LastName = auth.LastName
		authenticationUser.Email = auth.Email
		authenticationUser.Phone = auth.Phone
		authenticationUser.Role = auth.Role
		authenticationUser.TokenString = validToken
		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(authenticationUser)
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

		error := h.Store.UserStoreInterface.AddUser(user)

		if error != nil {
			http.Error(writer, error.Error(), http.StatusInternalServerError)
			return
		}

		validToken, err := helper.GenerateJWT(user.ID, user.Email, user.Role)
		if err != nil {
			http.Error(writer, "Failed to generate token", http.StatusInternalServerError)
			writer.Header().Set("Content-Type", "application/json")
			json.NewEncoder(writer).Encode(err)
			return
		}

		var authenticationUser structure.AuthUser

		authenticationUser.FirstName = user.FirstName
		authenticationUser.LastName = user.LastName
		authenticationUser.Email = user.Email
		authenticationUser.Phone = user.Phone
		authenticationUser.Role = user.Role
		authenticationUser.TokenString = validToken

		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(authenticationUser)
	}
}
