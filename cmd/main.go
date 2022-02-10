package main

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/tomiok/pelix-api/pkg/database"
	"github.com/tomiok/pelix-api/pkg/movies"
	"os"
	"runtime/trace"
)

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	err := godotenv.Load(".env")

	if err != nil {
		log.Error().Err(err)
	}

	migrate()
	//server := web.CreateServer()

	//go movies.ConcurrentJob()

	movies.Job(database.Get())
	//log.Fatal().Err(server.Run("8500"))
	//time.Sleep(1 * time.Hour)
}
