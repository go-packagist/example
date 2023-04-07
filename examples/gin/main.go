package main

import (
	"github.com/go-packagist/example/app/framework/provider"
	"github.com/go-packagist/example/app/http"
	"github.com/go-packagist/example/bootstrap"
)

func main() {
	app := bootstrap.App()
	app.Register(provider.NewGinProvider(app.Container))

	http.NewKernel(app).Boot()
}
