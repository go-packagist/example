package bootstrap

import (
	"github.com/go-packagist/example/app/framework/foundation"
	"github.com/go-packagist/example/app/framework/provider"
)

func App() *foundation.Application {
	app := foundation.NewApplication()

	app.Register(provider.NewDotenvProvider(app.Container))
	app.Register(provider.NewConfigProvider(app.Container))
	app.Register(provider.NewHashingProvider(app.Container))
	app.Register(provider.NewEncryptionProvider(app.Container))
	app.Register(provider.NewRediserProvider(app.Container))

	return app
}
