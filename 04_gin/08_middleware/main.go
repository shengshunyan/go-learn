package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	Name     string `json:"name" binding:"required,min=3,max=20"`
	Password string `json:"password" binding:"required"`
}

func MyLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		// 可以添加状态
		c.Set("name", "shane")

		// 让原来的逻辑继续执行
		//c.Next()
		c.Abort()

		duration := time.Since(start)
		fmt.Printf("took %s\n", duration)
		// 可以获取请求完成的状态码
		status := c.Writer.Status()
		fmt.Printf("status: %d\n", status)
	}
}

func main() {
	r := gin.Default()
	r.Use(MyLogger())
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
