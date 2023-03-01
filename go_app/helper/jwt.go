package helper

import (
	"back_project/structure"
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
	claims["id_user"] = id
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		return fmt.Errorf("something went wrong: %s", err.Error()).Error(), err
	}
	return tokenString, nil
}

func ExtractClaims(_ http.ResponseWriter, request *http.Request) (*structure.Token, error) {
	if request.Header["Token"] != nil {
		token, err := jwt.Parse(request.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("there was an error in parsing")
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			return nil, err
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			var currentToken structure.Token
			currentToken.Email = claims["email"].(string)
			currentToken.IDUser = int(claims["id_user"].(float64))
			currentToken.Role = claims["role"].(string)

			return &currentToken, nil
		}
	}

	return nil, fmt.Errorf("unable to extract claims")
}
