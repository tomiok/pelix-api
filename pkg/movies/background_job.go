package movies

import "time"

func BackgroundJob() {
	tick := time.NewTicker(1 * time.Second)

	for {
		select {
		case <-tick.C:

		}
	}
}
