package movies

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"os"
)

type MovieStorage interface {
	Lookup(title string) ([]MovieSearchRes, error)
	Upcoming() ([]MovieSearchRes, error)
}

type movieStorage struct {
	db *gorm.DB
}

func newMovieStorage(db *gorm.DB) *movieStorage {
	return &movieStorage{
		db: db,
	}
}

func (m *movieStorage) Lookup(title string) ([]MovieSearchRes, error) {
	fn := search()
	return fn(title)
}

func (m *movieStorage) Upcoming() ([]MovieSearchRes, error) {
	return upcoming()
}

func search() func(t string) ([]MovieSearchRes, error) {
	return func(t string) ([]MovieSearchRes, error) {
		uriFmt := "https://www.omdbapi.com/?s=%s&apikey=4ecb0111"
		res, err := http.Get(fmt.Sprintf(uriFmt, t))

		if err != nil {
			return nil, err
		}

		body := res.Body
		defer func() {
			_ = body.Close()
		}()

		var omdbRes OmdbResponse
		err = json.NewDecoder(body).Decode(&omdbRes)

		if err != nil {
			return nil, err
		}

		return omdbRes.ToApi(), nil
	}
}

func upcoming() ([]MovieSearchRes, error) {
	uriFmt := "https://api.themoviedb.org/3/movie/upcoming?api_key=%s"
	uri := fmt.Sprintf(uriFmt, os.Getenv("MOVIE_KEY"))

	res, err := http.Get(uri)

	if err != nil {
		return nil, err
	}

	var movies UpcomingMovies
	body := res.Body
	defer func() {
		_ = body.Close()
	}()
	err = json.NewDecoder(body).Decode(&movies)

	return movies.ToApi(), nil
}
