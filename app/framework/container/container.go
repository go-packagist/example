package container

import (
	"github.com/gin-gonic/gin"
	"github.com/go-packagist/encryption"
	"github.com/go-packagist/event/v2"
	"github.com/go-packagist/example/config"
	"github.com/go-packagist/hashing"
	"github.com/go-packagist/rediser/v2"
)

type Container struct {
	*config.Config
	Gin        *gin.Engine
	Hash       *hashing.Manager
	Encrypter  *encryption.Encrypter
	Rediser    *rediser.Manager
	Dispatcher *event.Dispatcher
}

func NewContainer() *Container {
	return &Container{}
}
