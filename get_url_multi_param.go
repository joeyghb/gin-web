package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func get_url_multi_param(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")

	c.JSON(http.StatusOK, gin.H{
		"name":   name,
		"action": action,
	})
}
