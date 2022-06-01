package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-packagist/example/bootstrap"
)

func main() {
	app := bootstrap.App()

	app.MustMake("gin").(*gin.Engine).Run(":8080")
}
