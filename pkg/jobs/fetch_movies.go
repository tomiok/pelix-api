package jobs

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tomiok/pelix-api/pkg/configs"
	"github.com/tomiok/pelix-api/pkg/movies"
	"net/http"
	"strings"
	"time"
)

/*
 ETL
	extract -> sacamos las peliculas de la base de datos de TMDB
	transform -> transformar el modelo de TMDB en nuestro modelo
	load -> guardarlo en nuestra base de datos

*/

const urlFmt = "https://api.themoviedb.org/3/movie/%d?api_key=%s"

var client = http.Client{
	Timeout: 3 * time.Second,
}

func Fetch(movieID uint) (*movies.TmdbByIdRes, error) {
	apiKey := configs.Get().MovieKey

	url := fmt.Sprintf(urlFmt, movieID, apiKey)

	res, err := client.Get(url)

	if err != nil {
		return nil, err
	}

	if !strings.Contains(res.Status, "200") {
		return nil, errors.New("status is not 200")
	}

	body := res.Body

	defer body.Close()

	var tmdbResult movies.TmdbByIdRes
	err = json.NewDecoder(body).Decode(&tmdbResult)

	if err != nil {
		return nil, err
	}

	return &tmdbResult, nil
}
