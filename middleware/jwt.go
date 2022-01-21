package middleware

import (
	"github.com/joho/godotenv"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func CreateTokens(userId uint, name string, role string) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err.Error())
	}
	claims := jwt.MapClaims{}

	claims["userId"] = userId
	claims["name"] = name
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 7).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	JWT := os.Getenv("JWT_SECRET")
	return token.SignedString([]byte(JWT))
}

func ExtractClaim(e echo.Context) (claims map[string]interface{}) {
	user := e.Get("user").(*jwt.Token)

	if user.Valid {
		claims = user.Claims.(jwt.MapClaims)
	}

	return
}
