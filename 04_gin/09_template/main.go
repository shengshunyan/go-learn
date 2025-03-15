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
	router := gin.Default()
	router.LoadHTMLGlob("/Users/shengshunyan/Desktop/go/go-learn/04_gin/09_template/templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	router.Run(":8080")
}
