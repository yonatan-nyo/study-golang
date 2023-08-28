package main

import (
	"example.com/m/v2/initializers"
	"example.com/m/v2/models"
)

func init() {
	initializers.LoadEnv()
	initializers.Database()
}
func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
