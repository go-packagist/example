package provider

import (
	"github.com/gin-gonic/gin"
	"github.com/go-packagist/example/app/framework/container"
)

type GinProvider struct {
	*container.Container
}

func NewGinProvider(container *container.Container) *GinProvider {
	return &GinProvider{
		Container: container,
	}
}

func (p *GinProvider) Register() {
	p.Gin = gin.Default()
}

func (p *GinProvider) Boot() {
}
