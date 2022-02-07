package movies

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"net/http"
	"os"
	"sync"
)

type MovieStorage interface {
	Lookup(title string) ([]MovieSearchRes, error)
	Upcoming() ([]MovieSearchRes, error)
	//SaveLocal save in local DB
	SaveLocal(movies []MovieSearchRes)
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
	result, err := upcoming()

	if err != nil {
		return nil, err
	}

	go m.SaveLocal(result)

	return result, err
}

func (m *movieStorage) SaveLocal(movies []MovieSearchRes) {
	var wg sync.WaitGroup
	for _, movie := range movies {
		wg.Add(1)
		go save(movie, m.db, &wg)
	}
	wg.Wait()
}

func save(m MovieSearchRes, db *gorm.DB, wg *sync.WaitGroup) {
	movie := m.ToMovie()

	err := db.Create(movie).Error

	if err != nil {
		log.Error().Err(err)
	}

	wg.Done()
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
