package config

import (
	db "github.com/go-packagist/framework/database"
	"github.com/go-packagist/framework/support/facades"
)

func database() {
	facades.MustConfig().Add("database", &db.Config{
		Default: "mysql",
		Connections: map[string]interface{}{
			"mysql": &db.MySqlConfig{
				Host:     "localhost",
				Port:     3306,
				Username: "root",
				Password: "123456",
				Database: "test",
			},
		},
	})
}
