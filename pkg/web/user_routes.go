package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tomiok/pelix-api/pkg/database"
	"github.com/tomiok/pelix-api/pkg/users"
)

func setupUserRoutes(app *fiber.App) {
	db := database.Get()
	web := users.NewWeb(db)

	group := app.Group("/users")

	group.Post("/", web.SaveHandler)
	group.Post("/login", web.LoginHandler)
}
