package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

type QueryParameters struct {
	Name string `form:"name"`
	Age  int    `form:"age"`
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")

	data := "Hello Go/Gin!!"

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{"data": data})
	})

	log.Println("server start at port 8080")
	log.Fatal(r.Run(":8080"))
}
