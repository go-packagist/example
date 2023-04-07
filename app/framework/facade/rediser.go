package facade

import (
	"github.com/go-packagist/example/app/framework/foundation"
	"github.com/go-packagist/rediser/v2"
)

func Rediser() *rediser.Manager {
	return foundation.GetInstance().Rediser
}
