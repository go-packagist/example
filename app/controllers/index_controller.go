package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-packagist/framework/database"
	"github.com/go-packagist/framework/foundation"
	"github.com/go-packagist/framework/support/facades"
	"gorm.io/gorm"
)

type indexController struct {
}

func NewIndexController() *indexController {
	return &indexController{}
}

func (c *indexController) Index(ctx *gin.Context) {
	type User struct {
		ID       int    `gorm:"primary_key"`
		Username string `gorm:"column:username"`
	}

	var user User
	foundation.App().MustMake("database").(*database.Manager).Connection().DB().(*gorm.DB).First(&user)

	ctx.JSON(200, gin.H{
		"message": "hello world~：" + facades.MustHash().Driver().MustMake("123456"),
		"user":    user,
	})
}
