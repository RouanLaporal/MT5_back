package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
)

func IsAuthorized(next func(writer http.ResponseWriter, request *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request.Header["Token"])
		if request.Header["Token"] != nil {
			token, err := jwt.Parse(request.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				_, ok := token.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					writer.WriteHeader(http.StatusUnauthorized)
					writer.Write([]byte("Unauthorized"))
				}
				return []byte(os.Getenv("JWT_SECRET")), nil
			})

			if err != nil {
				writer.WriteHeader(http.StatusUnauthorized)
				writer.Write([]byte("Unauthorized: " + err.Error()))
			}

			if token.Valid {
				next(writer, request)
			}
		} else {
			writer.WriteHeader(http.StatusUnauthorized)
			writer.Write([]byte("Unauthorized: not token found"))
		}
	})
}
