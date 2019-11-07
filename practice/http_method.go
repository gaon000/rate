package main

import (
	"strings"

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

	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		message := name + " is very handsome!"
		c.JSON(200, gin.H{"message": message})
	})

	r.GET("/user/:name/age/:old", func(c *gin.Context) {
		name := c.Param("name")
		age := c.Param("old")
		message := name + " is " + age + " years old."
		c.JSON(200, gin.H{"message": message})
	})

	r.GET("/colour/:colour/*fruits", func(c *gin.Context) {
		color := c.Param("colour")
		fruits := c.Param("fruits")
		fruitArray := strings.Split(fruits, "/")
		// remove the first element in fruit slice.
		fruitArray = append(fruitArray[:0], fruitArray[1:]...)
		c.JSON(200, gin.H{"color": color, "fruits": fruitArray})
	})
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
