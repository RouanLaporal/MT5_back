package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
)

func IsAuthorized(handler http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		if request.Header["Token"] == nil {
			http.Error(writer, "No Token Found", http.StatusUnauthorized)
			json.NewEncoder(writer).Encode("No Token Found")
			return
		}

		var mySigningKey = []byte(os.Getenv("JWT_SECRET"))

		token, err := jwt.Parse(request.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("there was an error in parsing")
			}
			return mySigningKey, nil
		})

		if err != nil {
			http.Error(writer, err.Error(), http.StatusUnauthorized)
			json.NewEncoder(writer).Encode(err)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if claims["role"] == "admin" {

				request.Header.Set("Role", "admin")
				handler.ServeHTTP(writer, request)
				return

			} else if claims["role"] == "user" {

				request.Header.Set("Role", "user")
				handler.ServeHTTP(writer, request)
				return
			}
		}
		http.Error(writer, "Not Authorized", http.StatusUnauthorized)
		json.NewEncoder(writer).Encode(err)
	}
}
