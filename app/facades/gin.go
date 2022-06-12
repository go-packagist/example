package facades

import (
	"github.com/gin-gonic/gin"
	"github.com/go-packagist/framework/foundation"
)

func Gin() (*gin.Engine, error) {
	g, err := foundation.App().Make("gin")

	if err != nil {
		return nil, err
	}

	return g.(*gin.Engine), nil
}

func MustGin() *gin.Engine {
	g, _ := Gin()

	return g
}
