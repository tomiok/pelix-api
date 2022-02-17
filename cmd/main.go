package main

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/tomiok/pelix-api/pkg/jobs"
	"github.com/tomiok/pelix-api/pkg/web"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	err := godotenv.Load(".env")

	if err != nil {
		log.Error().Err(err)
	}

	migrate()
	server := web.CreateServer()

	//go movies.ConcurrentJob()

	//job with thread pool
	//movies.Job(database.Get())


	//job as pipeline
	//movies.ETL()

	go jobs.ETL()
	log.Fatal().Err(server.Run("8500"))
	//time.Sleep(1 * time.Hour)
}
// https://discord.io/go-latam