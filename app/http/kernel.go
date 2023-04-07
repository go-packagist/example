package http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-packagist/example/app/framework/foundation"
	"github.com/go-packagist/example/app/http/routes"
)

var middlewares = []gin.HandlerFunc{
	gin.Logger(),
	gin.Recovery(),
}

type Kernel struct {
	*foundation.Application
}

func NewKernel(app *foundation.Application) *Kernel {
	return &Kernel{
		Application: app,
	}
}

func (k *Kernel) Boot() {
	k.Application.Boot()

	k.use(middlewares...)
	k.route()

	k.Gin.Run("localhost:1111")
}

func (k *Kernel) use(middlewares ...gin.HandlerFunc) {
	k.Gin.Use(middlewares...)
}

func (k *Kernel) route() {
	router := routes.NewRouter(k.Gin)
	router.Web()
}
