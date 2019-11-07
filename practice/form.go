package main

import (
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	r.POST("/form_post", formPost)

	return r
}

func formPost(c *gin.Context) {
	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "anonymous")

	headerType := c.GetHeader("Content-Type")

	c.JSON(200, gin.H{
		"status":              "posted",
		"message":             message,
		"nick":                nick,
		"header-content-type": headerType,
	})
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
