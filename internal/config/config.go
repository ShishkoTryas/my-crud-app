package config

import (
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	DB Postgres
}

type Postgres struct {
	Host     string `default:"localhost"`
	Port     int    `default:"5432"`
	Username string `default:"postgres"`
	DBName   string `default:"postgres"`
	SSLMode  string `default:"disable"`
	Password string `default:"02012001"`
}

func New() (*Config, error) {
	var cfg = new(Config)
	err := envconfig.Process("DB", &cfg.DB)
	if err != nil {
		log.Fatal("Can't create config", err)
	}

	return cfg, nil
}
