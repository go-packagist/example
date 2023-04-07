package provider

import (
	"github.com/go-packagist/encryption"
	"github.com/go-packagist/example/app/framework/container"
)

type EncryptionProvider struct {
	*container.Container
}

func NewEncryptionProvider(container *container.Container) *EncryptionProvider {
	return &EncryptionProvider{
		Container: container,
	}
}

func (p *EncryptionProvider) Register() {
	p.Encrypter = encryption.NewEncrypter(p.Config.App.Key)
}

func (p *EncryptionProvider) Boot() {

}
