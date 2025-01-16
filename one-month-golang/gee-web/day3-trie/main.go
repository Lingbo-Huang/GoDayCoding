package main

import (
	"gee"
	"net/http"
)

// 动态路由，即一条路由规则可以匹配某一类型而非某一条固定的路由。例如/hello/:name，可以匹配/hello/geektutu、hello/jack等

func main() {
	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>hello Gee</h1>")
	})

	r.GET("/hello", func(c *gee.Context) {
		c.String(http.StatusOK, "Hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{"filepath": c.Param("filepath")})
	})
	r.Run(":9999")
}
