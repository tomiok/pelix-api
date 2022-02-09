package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tomiok/pelix-api/pkg/database"
	"github.com/tomiok/pelix-api/pkg/movies"
)

func setUpMovieRoutes(app *fiber.App) {
	db := database.Get()
	web := movies.NewWeb(db)

	grp := app.Group("/movies")
	grp.Get("/", web.SearchMovieHandler)
	grp.Get("/upcoming", web.UpcomingMoviesHandler)

	grp.Post("/watchlist", web.AddMovieToListHandler)
	grp.Put("/watchlist", web.UpdateToSeenHandler)
	grp.Delete("/watchlist", web.DeleteMovieInListHandler)
}
