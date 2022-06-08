package config

import "github.com/go-packagist/framework/support/facades"

func redis() {
	facades.MustConfig().Add("redis", map[string]interface{}{
		"default": "redis",
		"connections": map[string]interface{}{
			"redis": map[string]interface{}{
				"driver":   "redis",
				"host":     "localhost",
				"port":     63790,
				"database": 0,
				"password": "",
			},
		},
	})
}
