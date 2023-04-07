package provider

import (
	"github.com/go-packagist/example/app/framework/container"
	"github.com/go-packagist/hashing"
)

type HashingProvider struct {
	*container.Container
}

func NewHashingProvider(container *container.Container) *HashingProvider {
	return &HashingProvider{
		Container: container,
	}
}

func (p *HashingProvider) Register() {
	p.Hash = hashing.NewManager(p.Config.Hash)
}

func (p *HashingProvider) Boot() {

}
