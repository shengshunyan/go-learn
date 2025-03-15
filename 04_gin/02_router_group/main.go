package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		v1.GET("/hen", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "ha",
			})
		})
	}

	v2 := r.Group("/api/v2")
	{
		v2.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		v2.GET("/hen", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "ha",
			})
		})
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
