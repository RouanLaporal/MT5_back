package web

import (
	"back_project/helper"
	"back_project/structure"
	"encoding/json"
	"net/http"
)

// User godoc
//
//	@Summary		UpdateUser
//	@Description	get user information by ID and modify this user
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			firstName		body	string	true	"user firstName"
//	@Param			lastName		body	string	true	"user lastName"
//	@Param			phone			body	string	true	"user phone"
//	@Param			email			body	string	true	"user email"
//	@Success		200
//	@Router			/user/update-profile [patch]
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

// User godoc
//
//		@Summary		VerifyPassword
//		@Description	Verify password of an user with its token, so we retrive the user with the email and password
//		@Tags			User
//		@Accept			json
//		@Produce		json
//	 	@Header			200		{string} 	Token	"user token"
//		@Param			email			body	string	true	"user email"
//		@Success		200
//		@Router			/user/verify-password [post]
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

// User godoc
//
//	@Summary		UpdatePassword
//	@Description	Update password of an user with its token, so we can change the password with the email
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Header			200		{string} 	Token	"user token"
//	@Success		200
//	@Router			/user/update-password [patch]
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

// User godoc
//
//	@Summary		DeleteUser
//	@Description	Delete a user from database with its id
//	@Tags			User
//	@Produce		json
//	@Header			200		{string} 	Token	"user token"
//	@Success		200
//	@Router			/user/delete-profile [delete]
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
