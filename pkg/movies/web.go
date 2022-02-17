package movies

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"net/http"
)

type Web struct {
	*Service
}

func NewWeb(db *gorm.DB) *Web {
	movieStorage := newMovieStorage(db)
	listStorage := newListStorage(db)
	service := NewService(movieStorage, listStorage)
	return &Web{
		service,
	}
}

func (w *Web) SearchMovieHandler(ctx *fiber.Ctx) error {
	title := ctx.Query("t", "fight club")
	res, err := w.Lookup(title)

	if err != nil {
		return err
	}

	return ctx.JSON(&res)
}

func (w *Web) UpcomingMoviesHandler(ctx *fiber.Ctx) error {
	res, err := w.Upcoming()

	if err != nil {
		return err
	}

	return ctx.JSON(res)
}

func (w *Web) HealthHandler(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"status": "OK"})
}

func (w *Web) AddMovieToListHandler(ctx *fiber.Ctx) error {
	var listItem ListItem
	err := ctx.BodyParser(&listItem)

	if err != nil {
		return err
	}

	err = w.Add(listItem)

	if err != nil {
		return err
	}
	ctx.Status(http.StatusCreated)
	return ctx.JSON(&ResponseOK{Message: "added ok"})
}

func (w *Web) UpdateToSeenHandler(ctx *fiber.Ctx) error {
	var u UpdateListRequest
	err := ctx.BodyParser(&u)

	if err != nil {
		return err
	}

	err = w.Update(u.UserID, u.MovieID)

	if err != nil {
		return err
	}
	ctx.Status(http.StatusNoContent)
	return ctx.JSON(ResponseOK{Message: "updated OK"})
}

func (w *Web) DeleteMovieInListHandler(ctx *fiber.Ctx) error {
	var u UpdateListRequest //delete request
	err := ctx.BodyParser(&u)

	if err != nil {
		return err
	}

	err = w.Delete(u.UserID, u.MovieID)

	if err != nil {
		return err
	}
	ctx.Status(http.StatusNoContent)
	return ctx.JSON(ResponseOK{Message: "updated OK"})
}

type ResponseOK struct {
	Message string `json:"msg"`
}
