package middlewares

import (
	"BelajarAPI/config"
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(hp string) (string, error) {
	var data = jwt.MapClaims{}
	data["hp"] = hp
	data["iat"] = time.Now().Unix()
	data["exp"] = time.Now().Add(time.Hour * 3).Unix()

	var processToken = jwt.NewWithClaims(jwt.SigningMethodHS256, data)

	result, err := processToken.SignedString([]byte(config.JWTSECRET))

	if err != nil {
		defer func () {
			if err := recover(); err != nil {
				log.Println("error jwt creation:", err)

			}
		}()
		return "", errors.New("terjadi masalah pembuatan")
	}

	return result, nil
}

func DecodeToken(token *jwt.Token) string {
	var result string
	var claim = token.Claims.(jwt.MapClaims)

	if val, found := claim["hp"]; found {
		result = val.(string)
	}

	return result
}