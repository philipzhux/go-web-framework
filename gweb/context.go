package gweb

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Context struct {
	W http.ResponseWriter
	Request *http.Request
	Route route_t
	StatusCode int
}
type AnyMap map[interface{}]interface{}

func (m AnyMap) MarshalJSON() ([]byte, error) {
	inter_map := map[string]interface{}{}
	for k,v := range(m) {
		inter_map[fmt.Sprintf("%v",k)] = v
	}
	return json.Marshal(inter_map)
}

func newContext(w http.ResponseWriter, request *http.Request) (*Context) {
	return &Context{
		W: w,
		Request: request,
		Route: route_t{Method: request.Method, Pattern: request.URL.Path},
	}
}
/* http.Request Gets */
func (c *Context) GetRequestPost(key string) (string) {
	return c.Request.FormValue(key)
}


func (c *Context) GetRequestQuery(key string) (string) {
	return c.Request.URL.Query().Get(key)
}


/* Response Set (Mainly Helper Functions) */
func (c *Context) SetResponseStatus(statusCode int){
	c.StatusCode = statusCode
	c.W.WriteHeader(statusCode)
}

func (c *Context) SetResponseHeader(key string, value string) {
	c.W.Header().Set(key,value)
}

/* Response Send (Main User Interfaces) */

func (c *Context) SendHttp(statusCode int, payload string) (error) {
	c.SetResponseStatus(statusCode)
	c.SetResponseHeader("Content-Type","text/html")
	_,err := c.W.Write([]byte(payload))
	return err
}

func (c *Context) SendJSON(statusCode int, payload interface{}) (error){
	c.SetResponseStatus(statusCode)
	c.SetResponseHeader("Content-Type","application/json")
	return json.NewEncoder(c.W).Encode(payload)
}

func (c *Context) SendString(statusCode int, format string, values ...interface{}) (error){
	c.SetResponseStatus(statusCode)
	c.SetResponseHeader("Content-Type","text/plain")
	_,err := fmt.Fprintf(c.W,format,values...)
	return err
}

func (c *Context) SendData(statusCode int, payload []byte) (error){
	c.SetResponseStatus(statusCode)
	_,err := c.W.Write(payload)
	return err
}

