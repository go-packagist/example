package config

import "github.com/go-packagist/framework/support/facades"

func init() {
	facades.MustConfig().Add("app", map[string]interface{}{
		"app_name": "Go Packagist",
	})
}
