package GolangWeb

import (
	"GolangWeb/log"
	"encoding/json"
	"fmt"
	"net/http"
)

type Context struct {
	Writer http.ResponseWriter

	Req *http.Request

	Status int

	Url string

	Method string

	Handlers []HandleFunc

	HandleFuncIndex int
}

func NewContext(writer http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer:          writer,
		Req:             req,
		Status:          http.StatusOK,
		Url:             req.URL.String(),
		Method:          req.Method,
		Handlers:        make([]HandleFunc, 0),
		HandleFuncIndex: -1,
	}
}

func (c *Context) GetValueFromPostForm(name string) interface{} {
	return c.Req.FormValue(name)
}

func (c *Context) QueryValueFromURL(name string) string {
	return c.Req.URL.Query().Get(name)
}

func (c *Context) SetStatus(status int) {
	c.Status = status
}

func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

func (c *Context) HTML(status int, html string) {
	c.SetStatus(status)
	c.SetHeader("Content-Type", "text/html")
	_, err := c.Writer.Write([]byte(html))
	if err != nil {
		log.Errorf("write html error: %v", err)
		return
	}
}

func (c *Context) STRING(status int, ftmStr string, args ...interface{}) {
	c.SetStatus(status)
	c.SetHeader("Content-Type", "text/plain")
	_, err := c.Writer.Write([]byte(fmt.Sprintf(ftmStr, args...)))
	if err != nil {
		log.Errorf("write string error: %v", err)
		return
	}
}

func (c *Context) JSON(status int, data interface{}) {
	c.SetStatus(status)
	c.SetHeader("Content-Type", "application/json")
	var err error
	if err == nil {
		err = json.NewEncoder(c.Writer).Encode(data)
	}
	if err != nil {
		log.Errorf("write json error: %v", err)
		http.Error(c.Writer, err.Error(), 500)
		return
	}
}

func (c *Context) Next() {
	c.HandleFuncIndex++
	length := len(c.Handlers)
	for c.HandleFuncIndex < length {
		c.Handlers[c.HandleFuncIndex](c)
		c.HandleFuncIndex++
	}
}
