package api_jwt

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/tomiok/pelix-api/pkg/configs"
	"net/http"
	"strings"
	"time"
)

func AuthMiddleware(ctx *fiber.Ctx) error {
	authHeader := ctx.Request().Header.Peek("Authorization")
	// Leaving just the token from the header because a header like
	// "Authorization: Bearer eyJhb..." is expected.
	authHeaderParts := strings.Split(string(authHeader), " ")

	if len(authHeader) > 0 && len(authHeaderParts) > 1 {
		token, _ := jwt.Parse(authHeaderParts[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(configs.Get().JwtSecret), nil
		})

		_, ok := token.Claims.(jwt.MapClaims)

		if ok && token.Valid {
			return ctx.Next()
		}
	}

	ctx.Status(http.StatusUnauthorized)
	return errors.New("invalid jwt")
}

func SignToken(id uint, email, username string) string {
	cfg := configs.Get()
	token := jwt.New(jwt.SigningMethodHS512)
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
