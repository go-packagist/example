package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-packagist/example/bootstrap"
	"github.com/go-packagist/example/routes"
	"github.com/go-packagist/framework/foundation"
)

func main() {
	bootstrap.App()

	routes.Routes()

	foundation.App().MustMake("gin").(*gin.Engine).Run(":8080")
}
