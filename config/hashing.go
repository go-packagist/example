package config

import "github.com/go-packagist/framework/support/facades"

func init() {
	facades.MustConfig().Add("hashing", map[string]interface{}{
		"driver": "md5",
	})
}
