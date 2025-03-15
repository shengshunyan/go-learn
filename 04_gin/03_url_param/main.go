package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/goods/:id", func(c *gin.Context) {
		// 获取路径参数
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"id": id,
		})
	})

	type Person struct {
		ID   string `uri:"id" binding:"required"`
		Name string `uri:"name"`
	}
	r.GET("/person/:id/:name", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBindUri(&person); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"id":   person.ID,
			"name": person.Name,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
