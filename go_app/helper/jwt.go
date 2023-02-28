package helper

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJWT(id int, email, role string) (string, error) {
	var mySigningKey = []byte(os.Getenv("JWT_SECRET"))
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["id"] = id
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		return fmt.Errorf("something went wrong: %s", err.Error()).Error(), err
	}
	return tokenString, nil
}

func ExtractClaims(_ http.ResponseWriter, request *http.Request) (string, error) {
	if request.Header["Token"] != nil {
		token, err := jwt.Parse(request.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("there was an error in parsing")
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			return "Error Parsing Token: ", err
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			email := claims["email"].(string)
			return email, nil
		}
	}

	return "unable to extract claims", nil
}
