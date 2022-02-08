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
	return ctx.JSON(&AddToListResponse{Message: "added ok"})
}

type AddToListResponse struct {
	Message string `json:"msg"`
}
