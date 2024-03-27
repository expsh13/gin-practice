package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type QueryParameters struct {
	Name string `form:"name"`
	Age  int    `form:"age"`
}

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, "Hello, Gin!")
	})
	r.GET("/article", func(c *gin.Context) {
		c.JSON(200, "postArticleHandler")
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 200 -> /user/john, 301 -> /user/john/, 404 -> /user/john/get
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// 200 -> /user/john/get, /user/john/get/ok
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	r.POST("/user/:name/*action", func(c *gin.Context) {
		// /user/john/get/ok も /user/:name/*actionと等価になる
		fmt.Println(c.FullPath() == "/user/:name/*action")
	})

	// welcome?firstname=Jane&lastname=Doe
	r.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	r.GET("/bind", func(c *gin.Context) {
		var queryParams QueryParameters
		// ShouldBindQueryを使ってクエリパラメータを構造体にバインド
		if err := c.ShouldBindQuery(&queryParams); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// バインドが成功したら、クエリパラメータを使った処理を行う
		c.JSON(http.StatusOK, gin.H{
			"name": queryParams.Name,
			"age":  queryParams.Age,
		})
	})

	r.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	v1 := r.Group("/v1")
	v1.GET("/get", func(c *gin.Context) {
		c.String(200, "/v1/GET")
	})

	v2 := r.Group("/v2")
	{
		v2.GET("/get", func(c *gin.Context) {
			c.String(200, "/v2/GET")
		})
	}

	// gin.H は map[string]interface{} へのショートカットです。
	r.GET("/someJSON", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("/moreJSON", func(c *gin.Context) {
		// 構造体を使うこともできます。
		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}
		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123
		// msg.Name は JSON 内で "user" となることに注意してください
		// 右記が出力されます  :   {"user": "Lena", "Message": "hey", "Number": 123}
		c.JSON(http.StatusOK, msg)
	})

	r.GET("/someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	log.Println("server start at port 8080")
	log.Fatal(r.Run(":8080"))
}
