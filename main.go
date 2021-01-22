package main

import (
	"net/http"

	"gating"
)

func main() {
	r := gating.New()
	r.GET("/index", func(c *gating.Context){
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gating.Context){
			c.HTML(http.StatusOK, "<h1>Hello, Gating</h1>")
		})
		v1.GET("/hello", func(c *gating.Context){
			// expect /hello?name=hwnzy
			c.String(http.StatusOK, "Hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *gating.Context){
			// expect /hello/hwnzy
			c.String(http.StatusOK, "Hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *gating.Context){
			c.JSON(http.StatusOK, gating.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})
	}

	r.Run(":8001")
}
