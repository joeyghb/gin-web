package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func get_url_by_query(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")

	firstname := c.DefaultQuery("firstname", "None")
	lastname  := c.Query("lastname")

	c.JSON(http.StatusOK, gin.H{
		"name":      name,
		"action":    action,
		"firstname": firstname,
		"lastname":  lastname,
	})
}