package facade

import (
	"github.com/go-packagist/example/app/framework/foundation"
	"github.com/go-packagist/hashing"
)

func Hash() *hashing.Manager {
	return foundation.GetInstance().Hash
}
