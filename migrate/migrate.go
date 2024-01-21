package main

import (
	"go-gin-crud/initializers"
	"go-gin-crud/models"
)

func init() {
	initializers.ConnectToDB()
	initializers.LoadEnv()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
