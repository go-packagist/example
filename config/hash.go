package config

import (
	"github.com/go-packagist/hashing"
)

var hash = &hashing.Config{
	Driver: "bcrypt",
}
