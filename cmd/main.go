package main

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
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

	log.Fatal().Err(server.Run("8500"))
}
