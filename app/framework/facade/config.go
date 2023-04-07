package facade

import (
	"github.com/go-packagist/example/app/framework/foundation"
	"github.com/go-packagist/example/config"
)

func Config() *config.Config {
	return foundation.GetInstance().Config
}
