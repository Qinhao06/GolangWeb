package GolangWeb

type RouterGroup struct {
	Engine *Engine

	Prefix string

	Parent *RouterGroup

	MiddleWares []HandleFunc
}

func (r *RouterGroup) Group(prefix string, middlewares ...HandleFunc) *RouterGroup {
	routerGroup := &RouterGroup{
		Parent:      r,
		Prefix:      r.Prefix + prefix,
		MiddleWares: middlewares,
		Engine:      r.Engine,
	}
	r.Engine.Groups = append(r.Engine.Groups, routerGroup)
	return routerGroup
}

func (r *RouterGroup) AddRouter(method string, url string, handleFunc HandleFunc) {
	r.Engine.AddRouter(method, r.Prefix+url, handleFunc)
}

func (r *RouterGroup) GET(url string, handleFunc HandleFunc) {
	r.AddRouter("GET", url, handleFunc)
}

func (r *RouterGroup) POST(url string, handleFunc HandleFunc) {
	r.AddRouter("POST", url, handleFunc)
}

func (r *RouterGroup) GetRouter(method string, url string) (HandleFunc, bool) {
	return r.Engine.GetRouter(method, r.Prefix+url)
}

func (r *RouterGroup) Run(addr string) error {
	return r.Engine.Run(addr)
}

func (r *RouterGroup) Use(middlewares ...HandleFunc) {
	r.MiddleWares = append(r.MiddleWares, middlewares...)
}
