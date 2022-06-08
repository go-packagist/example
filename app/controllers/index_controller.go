package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-packagist/framework/support/facades"
)

type indexController struct {
}

func NewIndexController() *indexController {
	return &indexController{}
}

func (c *indexController) Index(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "hello world~：" + facades.MustHash().Driver().MustMake("123456"),
	})
}
