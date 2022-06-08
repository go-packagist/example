package config

import "github.com/go-packagist/framework/support/facades"

func app() {
	facades.MustConfig().Add("app", map[string]interface{}{
		"name": "Go Packagist",
	})
}
