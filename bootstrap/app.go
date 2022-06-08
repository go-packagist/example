package bootstrap

import (
	c "github.com/go-packagist/example/config"
	"github.com/go-packagist/framework/config"
	"github.com/go-packagist/framework/foundation"
	"github.com/go-packagist/framework/hashing"
	"github.com/go-packagist/framework/redis"
	"github.com/go-packagist/gin"
)

func App() *foundation.Application {
	app := foundation.NewApplication("./")

	// configure
	app.Register(config.NewConfigProvider(app))
	c.Load()

	// register
	app.Register(hashing.NewHashProvider(app))
	app.Register(gin.NewGinProvider(app))
	app.Register(redis.NewRedisProvider(app))

	return app
}
