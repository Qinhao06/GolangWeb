package GolangWeb

import (
	"GolangWeb/log"
	"regexp"
	"strings"
)

type HandleFunc func(ctx *Context)

type Router struct {
	RouterMap      map[string]HandleFunc
	DynamicRouters []string
}

func NewRouter() *Router {
	return &Router{
		RouterMap:      make(map[string]HandleFunc),
		DynamicRouters: make([]string, 0),
	}
}

func (r *Router) AddRouter(method string, url string, handleFunc HandleFunc) {
	r.RouterMap[method+"-"+url] = handleFunc
	if r.isDynamicRouter(url) {
		r.DynamicRouters = append(r.DynamicRouters, url)
	}

}

func (r *Router) isDynamicRouter(url string) bool {
	return strings.Contains(url, "[") || strings.Contains(url, "*") ||
		strings.Contains(url, "\\") || strings.Contains(url, "+") ||
		strings.Contains(url, "?") || strings.Contains(url, "{") ||
		strings.Contains(url, "^") || strings.Contains(url, "$")
}

func (r *Router) GET(url string, handleFunc HandleFunc) {
	r.AddRouter("GET", url, handleFunc)
}

func (r *Router) POST(url string, handleFunc HandleFunc) {
	r.AddRouter("POST", url, handleFunc)
}

func (r *Router) GetRouter(method string, url string) (HandleFunc, bool) {
	handleFunc, ok := r.RouterMap[method+"-"+url]
	if ok {
		return handleFunc, ok
	}
	for _, dynamicRouter := range r.DynamicRouters {
		compile, err := regexp.Compile(dynamicRouter)
		if err != nil {
			log.Errorf("compile dynamic router %s is error", dynamicRouter)
			continue
		}
		if compile.MatchString(url) {
			return r.RouterMap[method+"-"+dynamicRouter], true
		}
	}
	return nil, false
}
