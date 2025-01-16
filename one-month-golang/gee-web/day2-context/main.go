package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.New()

	// curl -i http://localhost:9999/
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	// curl -i http://localhost:9999/hello\?name\=hlb
	r.GET("/hello", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	// curl -i "http://localhost:9999/login" -X POST -d 'username=hlb&password=1234'
	r.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	r.Run(":9999")
}
