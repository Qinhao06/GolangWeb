package GolangWeb

import (
	"GolangWeb/log"
	"net/http"
	"strings"
	"time"
)

type Engine struct {
	Router *Router

	*RouterGroup

	Groups []*RouterGroup
}

func (e *Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	context := NewContext(writer, request)
	var middlewares []HandleFunc
	for _, group := range e.Groups {
		if strings.HasPrefix(context.Req.URL.Path, group.Prefix) {
			middlewares = append(middlewares, group.MiddleWares...)
		}
	}
	context.Handlers = append(middlewares, e.MiddleWares...)
	err := e.handle(context)
	if err != nil {
		log.Errorf("handle error: %v", err)
	}

}

func (e *Engine) handle(context *Context) error {
	if handleFunc, ok := e.Router.GetRouter(context.Req.Method, context.Req.URL.Path); ok {
		context.Handlers = append(context.Handlers, handleFunc)
		context.Next()
	} else {
		_, err := context.Writer.Write([]byte("404 not found"))
		if err != nil {
			log.Errorf("write error: %v", err)
			return err
		}
	}
	return nil
}

func New() *Engine {
	e := &Engine{
		Router: NewRouter(),
	}
	e.Groups = make([]*RouterGroup, 0)
	e.RouterGroup = &RouterGroup{
		Engine: e,
	}
	e.Groups = append(e.Groups, e.RouterGroup)
	return e
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) AddRouter(method string, url string, handleFunc HandleFunc) {
	e.Router.AddRouter(method, url, handleFunc)
}

func (e *Engine) GET(url string, handleFunc HandleFunc) {
	e.Router.GET(url, handleFunc)
}

func (e *Engine) POST(url string, handleFunc HandleFunc) {
	e.Router.POST(url, handleFunc)
}

func (e *Engine) GetRouter(method string, url string) (HandleFunc, bool) {
	return e.Router.GetRouter(method, url)
}

func Default() *Engine {
	e := New()
	e.Use(Recovery(), Logger())
	return e
}

func Logger() HandleFunc {
	return func(context *Context) {
		t := time.Now()
		context.Next()
		log.Infof("[%d] %s in %v", context.Status, context.Req.RequestURI, time.Since(t))
	}
}
