package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "Hello, Gin!")
	})

	log.Println("server start at port 8080")
	log.Fatal(r.Run(":8080"))
}
