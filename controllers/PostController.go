package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func PostsGet(c *gin.Context) {
	c.JSON(200, gin.H{"message": "OK"})
}

func PostsCreate(c *gin.Context) {
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	fmt.Println(body)

	c.JSON(200, gin.H{"data": body})
}
