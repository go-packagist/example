package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-packagist/example/app/framework/facade"
)

type IndexController struct{}

func (c *IndexController) Index(ctx *gin.Context) {
	ciphertext, _ := facade.Encrypter().Encrypt("password")

	ctx.JSON(200, map[string]interface{}{
		"message": "Hello World",
		"hashing": facade.Hash().MustMake("password"),
		"encrypt": ciphertext,
	})
}

func (c *IndexController) Hello(ctx *gin.Context) {
	ctx.JSON(200, map[string]interface{}{
		"message": "Hello World",
		"rediser": facade.Rediser().Connect().Get(context.Background(), "key").Val(),
	})
}
