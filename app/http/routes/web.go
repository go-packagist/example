package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-packagist/example/app/http/controllers"
)

type Router struct {
	*gin.Engine
}

func NewRouter(e *gin.Engine) *Router {
	return &Router{e}
}

func (r *Router) Web() {
	indexController := new(controllers.IndexController)

	r.GET("/", indexController.Index)
	r.GET("/hello", indexController.Hello)
}
