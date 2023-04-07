package provider

import (
	"github.com/go-packagist/example/app/framework/container"
	"github.com/joho/godotenv"
)

type DotenvProvider struct {
	*container.Container
}

func NewDotenvProvider(container *container.Container) *DotenvProvider {
	return &DotenvProvider{
		Container: container,
	}
}

func (p *DotenvProvider) Register() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func (p *DotenvProvider) Boot() {

}
