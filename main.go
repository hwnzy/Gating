package main

import (
	"net/http"

	"gating"
)

func main() {
	r := gating.New()
	r.GET("/", func(c *gating.Context){
		c.HTML(http.StatusOK, "<h1>Hello Gating</h1>")
	})
	r.GET("/hello", func(c *gating.Context){
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "Hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *gating.Context){
		c.JSON(http.StatusOK, gating.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":8001")
}
