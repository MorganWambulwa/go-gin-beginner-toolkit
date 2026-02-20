package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the Go and Gin Beginner Toolkit",
		})
	})

	router.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "API is running",
		})
	})

	router.GET("/hello", func(c *gin.Context) {
		name := c.Query("name")
		if name == "" {
			name = "Stranger"
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, " + name + "!",
		})
	})

	router.GET("/ai", func(c *gin.Context) {
		prompt := c.Query("prompt")
		if prompt == "" {
			prompt = "Hello AI!"
		}
		response := "You asked: \"" + prompt + "\". This is a GenAI-style response."
		c.JSON(http.StatusOK, gin.H{
			"response": response,
		})
	})

	router.Run(":" + port)
}
