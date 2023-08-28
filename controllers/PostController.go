package controllers

import (
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
)

func PostsGet(c *gin.Context) {
	c.JSON(200, gin.H{"message": "OK"})
}

func PostsCreate(c *gin.Context) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(jsonData)

	c.JSON(200, gin.H{"data": jsonData})
}
