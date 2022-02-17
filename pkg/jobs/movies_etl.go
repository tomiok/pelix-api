package jobs

import (
	"github.com/rs/zerolog/log"
	"github.com/tomiok/pelix-api/pkg/database"
	"github.com/tomiok/pelix-api/pkg/movies"
	"gorm.io/gorm"
)

/*
 ETL
	extract -> sacamos las peliculas de la base de datos de TMDB
	transform -> transformar el modelo de TMDB en nuestro modelo
	load -> guardarlo en nuestra base de datos

	go routines
	channels
	pipeline

*/

func ETL() {
	db := database.Get()

	idsChannel := make(chan uint)
	tmdbChannel := make(chan *movies.TmdbByIdRes)
	moviesChannel := make(chan *movies.Movie)

	go extract(idsChannel, tmdbChannel)
	go transform(tmdbChannel, moviesChannel)
	go load(db, moviesChannel)

	for i := 670; i < 1000000; i++ {
		idsChannel <- uint(i)
	}

	close(idsChannel)
}

func extract(idsChannel chan uint, tmdbChannel chan *movies.TmdbByIdRes) {
	for id := range idsChannel {
		res, err := Fetch(id)

		if err != nil {
			log.Error().Msgf("cannot fetch %s", err.Error())
			//continue
		} else {
			tmdbChannel <- res
		}

		//algo de codigo aca

	}
	close(tmdbChannel)
}

func transform(tmdbChannel chan *movies.TmdbByIdRes, moviesChannel chan *movies.Movie) {
	for tmdb := range tmdbChannel {
		movie := tmdb.ToMovie()
		moviesChannel <- movie
	}
	close(moviesChannel)
}

func load(db *gorm.DB, moviesChannel chan *movies.Movie) {
	for movie := range moviesChannel {
		err := db.Create(movie).Error

		if err != nil {
			log.Error().Msgf("cannot save in DB %s", err.Error())
		}
	}
}
