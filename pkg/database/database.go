//go:build !test
// +build !test

package database

import (
	"github.com/tomiok/pelix-api/pkg/configs"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func initDB() *gorm.DB {
	var logLevel = logger.Info
	if configs.Get().Env != "local"{
		logLevel = logger.Silent
	}
	cfg := gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	}

	if DB == nil {
		db, err := gorm.Open(sqlite.Open("pelix.db"), &cfg)

		if err != nil {
			panic(err)
		}

		DB = db
		return DB
	}

	return DB
}

func Get() *gorm.DB {
	return initDB()
}
