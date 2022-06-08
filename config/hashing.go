package config

import "github.com/go-packagist/framework/support/facades"

func hashing() {
	facades.MustConfig().Add("hashing", map[string]interface{}{
		"driver": "md5",
	})
}
