package main

import (
	"github.com/go-packagist/example/app/gozero"
	"github.com/go-packagist/example/bootstrap"
)

func main() {
	app := bootstrap.App()

	gozero.NewKernel(app).Boot()
}
