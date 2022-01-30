package configs

import "os"

var Cfg *Config

type Config struct {
	JwtSecret string
}

func Get() *Config {
	if Cfg == nil {
		Cfg = &Config{
			JwtSecret: os.Getenv("JWT_SECRET"),
		}
	}

	return Cfg
}
