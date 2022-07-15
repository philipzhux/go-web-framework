package gweb

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Context struct {
	w http.ResponseWriter;
	request *http.Request;
	method string;
	path string;
	statusCode int;
	parmas map[string]string;
	pattern string;
}
type AnyMap map[interface{}]interface{}

func (m AnyMap) MarshalJSON() ([]byte, error) {
	inter_map := map[string]interface{}{}
	for k,v := range(m) {
		inter_map[fmt.Sprintf("%v",k)] = v
	}
	return json.Marshal(inter_map)
}

func newContext(w http.ResponseWriter, request *http.Request, e *Engine) (*Context) {
	pattern,params := e.router.retrieve(request.Method,request.URL.Path)
	return &Context{
		w: w,
		request: request,
		method: request.Method,
		path: request.URL.Path,
		parmas: params,
		pattern: pattern,
	}
}
/* Getters */

func (c *Context) GetRequest() (*http.Request) {
	return c.request
}

func (c *Context) GetRequestPost(key string) (string) {
	return c.request.FormValue(key)
}

func (c *Context) GetParmas(key string) (string,bool) {
	parma, ok := c.parmas[key]
	return parma,ok
}

func (c *Context) GetRequestQuery(key string) (string) {
	return c.request.URL.Query().Get(key)
}

func (c *Context) GetPath() (string) {
	return c.path
}

func (c *Context) GetMethod() (string) {
	return c.method
}

func (c *Context) GetWriter() (http.ResponseWriter) {
	return c.w
}

func (c *Context) GetStatusCode() (int) {
	return c.statusCode
}

/* Setters */
/* Response Set (Mainly Helper Functions) */
func (c *Context) SetResponseStatus(statusCode int){
	c.statusCode = statusCode
	c.w.WriteHeader(statusCode)
}

func (c *Context) SetResponseHeader(key string, value string) {
	c.w.Header().Set(key,value)
}

/* Response Send (Main User Interfaces) */

func (c *Context) SendHttp(statusCode int, payload string) (error) {
	c.SetResponseStatus(statusCode)
	c.SetResponseHeader("Content-Type","text/html")
	_,err := c.w.Write([]byte(payload))
	return err
}

func (c *Context) SendJSON(statusCode int, payload interface{}) (error){
	c.SetResponseStatus(statusCode)
	c.SetResponseHeader("Content-Type","application/json")
	return json.NewEncoder(c.w).Encode(payload)
}

func (c *Context) SendString(statusCode int, format string, values ...interface{}) (error){
	c.SetResponseStatus(statusCode)
	c.SetResponseHeader("Content-Type","text/plain")
	_,err := fmt.Fprintf(c.w,format,values...)
	return err
}

func (c *Context) SendData(statusCode int, payload []byte) (error){
	c.SetResponseStatus(statusCode)
	_,err := c.w.Write(payload)
	return err
}

