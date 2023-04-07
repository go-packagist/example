package facade

import (
	"github.com/go-packagist/encryption"
	"github.com/go-packagist/example/app/framework/foundation"
)

func Encrypter() *encryption.Encrypter {
	return foundation.GetInstance().Encrypter
}
