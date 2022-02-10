package movies

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/tomiok/pelix-api/pkg/configs"
	"gorm.io/gorm"
	"net/http"
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
	mainChan := make(chan uint)

	for i := 0; i < threads; i++ {
		go job(mainChan, db)
	}

	for i := 1; i < maxRun; i++ {
		time.Sleep(700 * time.Millisecond)
		mainChan <- uint(i)
	}

	close(mainChan)
}

func job(idChan chan uint, db *gorm.DB) {
	for id := range idChan {
		res, err := FetchMovie(id)

		if err != nil {
			log.Error().Err(err)
			return
		}

		err = saveLocal(db, res)

		if res != nil {
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

func FetchMovie(movieID uint) (*TmdbByIdRes, error) {
	url := fmt.Sprintf(urlFmt, movieID, configs.Get().MovieKey)
	res, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	body := res.Body

	defer func() {
		_ = body.Close()
	}()
	var result TmdbByIdRes
	err = json.NewDecoder(body).Decode(&result)

	return &result, nil
}
