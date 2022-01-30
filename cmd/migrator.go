package main

import (
	"github.com/rs/zerolog/log"
	"github.com/tomiok/pelix-api/pkg/database"
	"github.com/tomiok/pelix-api/pkg/users"
)

func migrate() {
	db := database.Get()
	if err := users.MigrateModels(db); err != nil {
		log.Error().Msgf("cannot migrate users: %s", err.Error())
	}
}
