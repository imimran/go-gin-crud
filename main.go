package main

import (
	"fmt"
	"go-gin-crud/controllers"
	"go-gin-crud/initializers"
	"go-gin-crud/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	//Serving static files
	r.Static("/assets", "./assets")

	// Allow all Origins
	r.Use(cors.Default())

	r.Use(middlewares.ErrorHandler())

	r.POST("/post", controllers.PostCreate)
	r.GET("/post", middlewares.RequireAuth, controllers.AllPost)
	r.GET("/post/:id", controllers.GetPost)
	r.PUT("/post/:id", controllers.UpdatePost)
	r.DELETE("/post/:id", controllers.RemovePost)

	r.POST("/register", controllers.SignUp)
	r.POST("/login", controllers.Login)

	fmt.Println("Hello")
	r.Run() // listen and serve on 0.0.0.0:8080
}
