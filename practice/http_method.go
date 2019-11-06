package main

import (
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	r.GET("/someGet", someMethod)
	r.POST("/somePost", someMethod)
	r.PUT("/somePut", someMethod)
	r.DELETE("/someDelete", someMethod)
	r.PATCH("/somePatch", someMethod)
	r.HEAD("/someHead", someMethod)
	r.OPTIONS("/someOptions", someMethod)

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}

func someMethod(c *gin.Context) {
	httpMethod := c.Request.Method
	c.JSON(200, gin.H{"status": "good", "sending": httpMethod})
}
