package test

import (
	"GolangWeb"
	"fmt"
	"net/http"
	"testing"
)

func Test(t *testing.T) {
	r := GolangWeb.New()
	r.Router.GET("/", func(c *GolangWeb.Context) {
		fmt.Fprintf(c.Writer, "URL.Path = %q\n", c.Req.URL.Path)
	})

	r.Router.GET("/hello", func(ctx *GolangWeb.Context) {
		for k, v := range ctx.Req.Header {
			fmt.Fprintf(ctx.Writer, "Header[%q] = %q\n", k, v)
		}
	})

	r.Run(":9999")
}

func Test2(t *testing.T) {
	r := GolangWeb.New()
	r.Router.GET("/", func(c *GolangWeb.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	r.Router.GET("/hello", func(c *GolangWeb.Context) {
		// expect /hello?name=geektutu
		c.STRING(http.StatusOK, "hello %s, you're at %s\n", c.QueryValueFromURL("name"), c.Url)
	})

	r.Router.POST("/login", func(c *GolangWeb.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"username": c.GetValueFromPostForm("username"),
			"password": c.GetValueFromPostForm("password"),
		})
	})

	r.Run(":9999")
}

func Test22(t *testing.T) {
	r := GolangWeb.Default()
	r.GET("/", func(c *GolangWeb.Context) {
		c.STRING(http.StatusOK, "Hello \n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *GolangWeb.Context) {
		names := []string{"hello"}
		c.STRING(http.StatusOK, names[100])
	})

	r.Run(":9999")
}
