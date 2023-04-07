package config

import (
	"github.com/go-packagist/rediser/v2"
	"github.com/redis/go-redis/v9"
)

var rediserConf = &rediser.Config{
	Default: "default",
	Connections: map[string]rediser.Configable{
		"default": &redis.Options{
			Addr: "localhost:6379",
			DB:   0,
		},
	},
}
