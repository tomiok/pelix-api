package movies

import (
	"github.com/rs/zerolog/log"
	"github.com/tomiok/pelix-api/pkg/database"
	"time"
)

func BackgroundJob() {
	tick := time.NewTicker(1 * time.Second)

	for {
		select {
		case <-tick.C:

		}
	}
}

func ConcurrentJob() {
	log.Info().Msg("started!!!")
	mainCh := make(chan uint, 1000)

	for i := 0; i < 8; i++ {
		go selectSQL(mainCh)
	}

	for i := 0; i < 500; i++ {
		mainCh <- uint(i)
	}
	close(mainCh)
}

func normalizeJob(inputSql chan uint) {
	selectSQL(inputSql)
}

func selectSQL(inputSql chan uint) {
	for id := range inputSql {
		var movie Movie
		log.Info().Msgf("%d", database.Get().Where("id =?", id).First(&movie).RowsAffected)
	}
}
