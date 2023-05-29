package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	AppPort string `envconfig:"APP_PORT" default:"8080"`

	DBHost     string `envconfig:"DB_HOST" default:"127.0.0.1"`
	DBPort     string `envconfig:"DB_PORT" default:"3306"`
	DBUsername string `envconfig:"DB_USERNAME" default:"root"`
	DBPassword string `envconfig:"DB_PASSWORD" default:""`
	DBName     string `envconfig:"DB_NAME" default:"appointment"`
	DBDebug    bool   `envconfig:"DB_Debug" default:"true"`
}

func New() Config {
	var c Config
	err := envconfig.Process("", &c)
	if err != nil {
		log.Fatal(err.Error())
	}

	return c
}
