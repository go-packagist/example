package facade

import (
	"github.com/gin-gonic/gin"
	"github.com/go-packagist/example/app/framework/foundation"
)

func Gin() *gin.Engine {
	return foundation.GetInstance().Gin
}
