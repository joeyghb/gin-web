package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func get_url_one_param(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"message": "hello " + name,
	})
}
