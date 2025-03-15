package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	Name     string `json:"name" binding:"required,min=3,max=20"`
	Password string `json:"password" binding:"required"`
}

func main() {
	r := gin.Default()
	r.POST("/login", func(c *gin.Context) {
		var loginForm LoginForm
		if err := c.ShouldBind(&loginForm); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "login success",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
