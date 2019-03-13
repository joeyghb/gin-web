package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main_index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Main website",
		"message": "Hello Go Gin!!",
	})
}

func posts_index(c *gin.Context) {
	c.HTML(http.StatusOK, "posts/posts.tmpl", gin.H{
		"title": "Posts",
	})
}

func users_index(c *gin.Context) {
	c.HTML(http.StatusOK, "users/users.tmpl", gin.H{
		"title": "Users",
	})
}