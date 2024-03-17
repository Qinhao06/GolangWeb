package test

import (
	"GolangWeb"
	"log"
	"net/http"
	"testing"
	"time"
)

func onlyForV2() GolangWeb.HandleFunc {
	return func(c *GolangWeb.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.Status, c.Req.RequestURI, time.Since(t))
	}
}

func Test1(t *testing.T) {
	r := GolangWeb.New()
	r.Use(Logger()) // global midlleware
	r.GET("/", func(c *GolangWeb.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2()) // v2 group middleware
	v2.GET("/hello", func(c *GolangWeb.Context) {
		// expect /hello/geektutu
		c.STRING(http.StatusOK, "hello %s, you're at %s\n", c.QueryValueFromURL("name"), c.Url)
	})

	r.Run(":9999")
}

func Logger() GolangWeb.HandleFunc {
	return func(c *GolangWeb.Context) {
		// Start timer
		t := time.Now()
		// Process request
		// Calculate resolution time
		log.Printf("[%d] %s in %v", c.Status, c.Req.RequestURI, time.Since(t))
	}
}
