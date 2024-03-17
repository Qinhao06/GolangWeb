package test

import (
	"GolangWeb"
	"fmt"
	"testing"
)

func a(ctx *GolangWeb.Context) {
	fmt.Println("hello")
}

func newTestRouter() *GolangWeb.Router {
	r := GolangWeb.NewRouter()
	r.AddRouter("GET", "/", a)
	r.AddRouter("GET", "/hello/*", a)
	r.AddRouter("GET", "/hello/b/c", a)
	r.AddRouter("GET", "/hi/[0-9]", a)
	r.AddRouter("GET", "/assets/*filepath", a)
	return r
}

func TestGetRoute(t *testing.T) {
	r := newTestRouter()
	n, ps := r.GetRouter("GET", "/hi/1")

	if n == nil {
		t.Fatal("nil shouldn't be returned")
	}

	fmt.Printf("matched path: , params['name']: %t\n", ps)

}
