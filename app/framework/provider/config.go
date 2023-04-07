package provider

import (
	"github.com/go-packagist/example/app/framework/container"
	"github.com/go-packagist/example/config"
)

type ConfigProvider struct {
	*container.Container
}

func NewConfigProvider(container *container.Container) *ConfigProvider {
	return &ConfigProvider{
		Container: container,
	}
}

func (p *ConfigProvider) Register() {
	p.Config = config.Load()
}

func (p *ConfigProvider) Boot() {
}
