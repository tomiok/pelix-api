package api_jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/tomiok/pelix-api/pkg/configs"
	"time"
)

func SignToken(id uint, email, username string) string {
	cfg := configs.Get()
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	claims["sub"] = id
	claims["email"] = email
	claims["username"] = username

	t, err := token.SignedString([]byte(cfg.JwtSecret))

	if err != nil {
		return ""
	}

	return t
}
