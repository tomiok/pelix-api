package movies

import (
	"github.com/rs/zerolog/log"
	"time"
)

func NormalizeJob() { //TODO finish this as a real normalization job
	ticker := time.NewTicker(1 * time.Second)

	for {
		select {
		case <- ticker.C:
			log.Info().Msg("done!")
		}
	}

}
