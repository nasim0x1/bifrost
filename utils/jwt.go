package utils

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/nasim0x1/bifrost/configs"
	"github.com/nasim0x1/bifrost/handlers"
)

func CreateJwtToken(userID int) (string, error) {
	expiration := time.Second * time.Duration(configs.Envs.JWTExpirationInSeconds)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    strconv.Itoa(int(userID)),
		"expiresAt": time.Now().Add(expiration).Unix(),
	})
	tokenString, err := token.SignedString(configs.Envs.GetJwtSecret())
	if err != nil {
		return "", err
	}
	return tokenString, err
}

func WithJwtAuth(handleFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := GetTokenFromRequest(r)
		token, err := validateJWT(tokenString)
		if err != nil {
			log.Println("invalid token err")
			handlers.SendErrorResponse(w, http.StatusUnauthorized, "Invalid token")
			return
		}
		if !token.Valid {
			log.Println("invalid token")
			handlers.SendErrorResponse(w, http.StatusUnauthorized, "Invalid token")
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		str := claims["userID"].(string)

		userID, err := strconv.Atoi(str)
		if err != nil {
			log.Printf("failed to convert userID to int: %v", err)
			handlers.SendErrorResponse(w, http.StatusUnauthorized, "Invalid token")

			return
		}
		log.Printf("userID: %v", userID)

		handleFunc(w, r)
	}
}

func validateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(configs.Envs.GetJwtSecret()), nil
	})
}
