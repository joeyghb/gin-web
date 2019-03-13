package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	// Ping test
	c.JSON(http.StatusOK, gin.H{
		"pong": "Hello",
	})
}