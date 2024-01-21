package controllers

import (
	"go-gin-crud/initializers"
	"go-gin-crud/models"

	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {

	//req.body
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	// create post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(201, gin.H{
		"data": post,
	})

}

func AllPost(c *gin.Context) {

	var posts []models.Post

	initializers.DB.Find(&posts)
	c.JSON(201, gin.H{
		"data": posts,
	})
}

func GetPost(c *gin.Context) {

	id := c.Param("id")

	var post models.Post

	initializers.DB.First(&post, id)
	c.JSON(201, gin.H{
		"data": post,
	})
}

func UpdatePost(c *gin.Context) {

	// id
	id := c.Param("id")

	//req.body
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	var post models.Post

	initializers.DB.First(&post, id)

	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	c.JSON(201, gin.H{
		"data": post,
	})
}

func RemovePost(c *gin.Context) {

	// id
	id := c.Param("id")

	var post models.Post
	initializers.DB.First(&post, id)

	initializers.DB.Delete(&models.Post{}, id)

	c.JSON(201, gin.H{
		"data": "Delete",
	})
}
