package configs

import "os"

var Cfg *Config

type Config struct {
	JwtSecret string
	MovieKey  string
	Env       string
}

func Get() *Config {
	if Cfg == nil {
		Cfg = &Config{
			JwtSecret: os.Getenv("JWT_SECRET"),
			MovieKey:  os.Getenv("MOVIE_KEY"),
			Env:       os.Getenv("ENV"),
		}
	}

	return Cfg
}
