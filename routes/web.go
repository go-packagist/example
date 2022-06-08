package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-packagist/example/app/controllers"
	"github.com/go-packagist/framework/foundation"
)

func Routes() {
	index := controllers.NewIndexController()

	foundation.App().MustMake("gin").(*gin.Engine).GET("/", index.Index)
}
