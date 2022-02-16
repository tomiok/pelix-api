package movies

import "github.com/tomiok/pelix-api/pkg/database"

func ETL() {
	idsChannel := make(chan uint)
	tmdbChan := make(chan *TmdbByIdRes)
	moviesChan := make(chan *Movie)
	db := database.Get()

	go extract(idsChannel, tmdbChan)
	go transform(tmdbChan, moviesChan)
	go load(moviesChan, db)

	for i := 0; i < 5000; i++ {
		idsChannel <- uint(i)
	}
	close(idsChannel)
}