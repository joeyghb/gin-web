package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func get_post_form(c *gin.Context) {
	// 获取post过来的message内容
	// 获取的所有参数内容的类型都是 string
	message := c.PostForm("message")

	// 如果不存在，使用第二个当做默认内容
	nick := c.DefaultPostForm("nick", "anonymous")

	c.JSON(http.StatusOK, gin.H{
		"status":  "posted",
		"message": message,
		"nick":    nick,
	})
}