package provider

import (
	"github.com/go-packagist/example/app/framework/container"
	"github.com/go-packagist/rediser/v2"
)

type RediserProvider struct {
	*container.Container
}

func NewRediserProvider(container *container.Container) *RediserProvider {
	return &RediserProvider{
		Container: container,
	}
}

func (p *RediserProvider) Register() {
	p.Rediser = rediser.New(p.Config.Rediser)
}

func (p *RediserProvider) Boot() {

}
