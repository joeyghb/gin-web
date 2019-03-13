package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


//curl -v -X POST http://localhost:18080/post -H 'content-type: application/json' -d '{ "user": "Pgluffy" }'

func get_post_json_format(c *gin.Context) {
	// get raw data
	d, err := c.GetRawData()
	if err!=nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": "get post data error!!",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": string(d),
	})

}