package movies

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/tomiok/pelix-api/pkg/configs"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

const urlFmt = "https://api.themoviedb.org/3/movie/%d?api_key=%s"

type TmdbByIdRes struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	ReleaseDate string `json:"release_date"`
	PosterPath  string `json:"poster_path"`
	Overview    string `json:"overview"`
	Runtime     int    `json:"runtime"`
	ImdbID      string `json:"imdb_id"`
}

func (r *TmdbByIdRes) ToMovie() *Movie {
	return &Movie{
		Title:    r.Title,
		Year:     r.ReleaseDate,
		ImdbID:   &r.ImdbID,
		TmdbID:   &r.ID,
		Runtime:  fmt.Sprintf("%d", r.Runtime),
		Released: r.ReleaseDate,
		Poster:   r.PosterPath,
		OverView: r.Overview,
	}
}

const maxRun = 1000000
const threads = 8

func Job(db *gorm.DB) {
	inputChan := make(chan uint, 1000)

	for i := 0; i < threads; i++ {
		go func() {
			job(inputChan, db)
		}()
	}

	for i := 1; i < maxRun; i++ {
		log.Info().Msgf("run #%d", i)
		inputChan <- uint(i)
	}

	close(inputChan)
}

func job(inputChan chan uint, db *gorm.DB) {
	for id := range inputChan {
		log.Info().Msgf("fetching %d", id)
		res, err := FetchMovie(id)

		if err != nil {
			log.Error().Msg(err.Error())
			continue
		}

		err = saveLocal(db, res)

		if err != nil {
			log.Error().Err(err)
		} else {
			log.Info().Msg("movie saved OK in local")
		}
	}
}

func saveLocal(db *gorm.DB, res *TmdbByIdRes) error {
	movie := res.ToMovie()
	return db.Create(movie).Error
}

func transport() *http.Transport {
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = 100
	t.MaxConnsPerHost = 100
	t.MaxIdleConnsPerHost = 100
	return t
}

var client = &http.Client{
	Timeout:   10 * time.Second,
	Transport: transport(),
}

func FetchMovie(movieID uint) (*TmdbByIdRes, error) {
	url := fmt.Sprintf(urlFmt, movieID, configs.Get().MovieKey)

	res, err := client.Get(url)

	if err != nil {
		return nil, err
	}

	if res.Status != "200 OK" {
		client.CloseIdleConnections()
		return nil, errors.New("wrong status " + res.Status + " for id " + strconv.Itoa(int(movieID)))
	}

	body := res.Body

	defer func() {
		_ = body.Close()
	}()
	var result TmdbByIdRes
	err = json.NewDecoder(body).Decode(&result)

	return &result, nil
}
