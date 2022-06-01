package bootstrap

import (
	"github.com/go-packagist/framework/config"
	"github.com/go-packagist/framework/foundation"
	"github.com/go-packagist/framework/hashing"
	"github.com/go-packagist/framework/support/facades"
)

func App() *foundation.Application {
	app := foundation.NewApplication("./")

	app.Register(config.NewConfigProvider(app))

	facades.MustConfig().Add("app", map[string]interface{}{
		"app_name": "Go Packagist",
	})

	facades.MustConfig().Add("hashing", map[string]interface{}{
		"driver": "md5",
	})

	app.Register(hashing.NewHashProvider(app))

	return app
}
