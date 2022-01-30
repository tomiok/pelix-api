package users

import (
	"github.com/gofiber/fiber/v2"
	api_jwt "github.com/tomiok/pelix-api/pkg/jwt"
	"gorm.io/gorm"
	"net/http"
)

type Web struct {
	*UserService
}

func NewWeb(db *gorm.DB) *Web {
	storage := newUserStorage(db)
	userService := NewService(storage)
	return &Web{
		UserService: userService,
	}
}

func (w *Web) SaveHandler(ctx *fiber.Ctx) error {
	var dto UserCreateDTO
	err := ctx.BodyParser(&dto)

	if err != nil {
		return err
	}

	u, err := w.Save(&dto)

	if err != nil {
		return err
	}

	return ctx.JSON(&UserResponseCreate{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
	})
}

func (w *Web) LoginHandler(ctx *fiber.Ctx) error {
	var dto LoginDTO
	err := ctx.BodyParser(&dto)

	if err != nil {
		return err
	}
	u, err := w.Login(dto)

	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return err
	}

	res := loginOK(u)

	return ctx.JSON(&res)
}

type UserResponseCreate struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type LoginResponse struct {
	Jwt string `json:"jwt"`
}

func loginOK(u *User) LoginResponse {
	token := api_jwt.SignToken(u.ID, u.Email, u.Username)
	return LoginResponse{
		Jwt: token,
	}
}
