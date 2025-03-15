package main

import (
	"github.com/gin-gonic/gin"
	"go-learn/04_gin/06_protobuf/proto"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/protobuf", func(c *gin.Context) {

		person := proto.Person{
			Name: "shane",
			Age:  20,
		}

		c.ProtoBuf(http.StatusOK, &person)
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
