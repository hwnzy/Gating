package main

import (
	"net/http"
	"html/template"
	"fmt"
	"log"
	"time"

	"gating"
)

type student struct {
	Name string
	Age int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := gating.New()
	r.Use(gating.Logger())
	
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "/home/hwnzy/Gating/static")
	// r.GET("/index", func(c *gating.Context){
	// 	c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	// })
	r.GET("/css", func(c *gating.Context) {
		c.HTML(http.StatusOK, "css.html", nil)
	})
	r.GET("/students", func(c *gating.Context) {
		c.HTML(http.StatusOK, "arr.html", gating.H{
			"title": "gating",
			"stuArr": [2]*student{&student{Name: "ningzhiying", Age: 26}, &student{Name: "huangting", Age: 25}},
		})
	})
	r.GET("/date", func(c *gating.Context) {
		c.HTML(http.StatusOK, "custom_func.html", gating.H{
			"title": "gating",
			"now":	time.Date(2021, 1, 24, 0, 0, 0, 0, time.UTC),
		})
	})

	v1 := r.Group("/v1")
	{
		// v1.GET("/", func(c *gating.Context){
		// 	c.HTML(http.StatusOK, "<h1>Hello, Gating</h1>")
		// })
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
