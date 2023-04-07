package foundation

import (
	"github.com/go-packagist/example/app/framework/container"
	"github.com/go-packagist/example/app/framework/provider"
)

var instance *Application

type Application struct {
	*container.Container
	providers []provider.Provider
	booted    bool
}

func NewApplication() *Application {
	app := &Application{
		Container: container.NewContainer(),
		providers: make([]provider.Provider, 10),
	}

	SetInstance(app)

	return app
}

func SetInstance(app *Application) {
	instance = app
}

func GetInstance() *Application {
	return instance
}

func (a *Application) Register(provider provider.Provider) {
	provider.Register()

	a.providers = append(a.providers, provider)

	if a.booted {
		provider.Boot()
	}
}

func (a *Application) Boot() {
	if a.booted {
		return
	}

	for _, p := range a.providers {
		if p != nil {
			p.Boot()
		}
	}

	a.booted = true
}
