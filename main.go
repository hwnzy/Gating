package main

import (
	"net/http"
	"log"
	"time"

	"gating"
)

func main() {
	r := gating.New()
	r.Use(gating.Logger())
	r.Static("/assets", "/home/hwnzy/Gating/static")
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
	v2.Use(func(c *gating.Context) {
		t := time.Now()
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	})
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
