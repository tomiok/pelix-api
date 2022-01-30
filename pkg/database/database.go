// +build !test

package database

import (
	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func initDB() *gorm.DB {
	if DB == nil {
		db, err := gorm.Open(sqlite.Open("pelix.db"))

		if err != nil {
			panic(err)
		}

		log.Debug().Msg("database crated")
		DB = db
		return DB
	}

	return DB
}

func Get() *gorm.DB {
	return initDB()
}

