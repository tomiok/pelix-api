package movies

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Web struct {
	*Service
}

func NewWeb(db *gorm.DB) *Web {
	storage := newMovieStorage(db)
	service := NewService(storage)
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
