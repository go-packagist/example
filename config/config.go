package config

import (
	"github.com/go-packagist/hashing"
	"github.com/go-packagist/rediser/v2"
)

type Config struct {
	App     *App
	Hash    *hashing.Config
	Rediser *rediser.Config
	Gozero  *Gozero
}

func Load() *Config {
	return &Config{
		App:     app,
		Hash:    hash,
		Rediser: rediserConf,
		Gozero:  gozero,
	}
}
