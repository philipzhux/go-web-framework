package gweb

import (
	"fmt"
	"log"
	"strings"
)

type route_t struct{
	method string;
	pattern string
}

type router_t struct {
	handlers map[route_t]handleFunc;
	tries map[string]*node
}

func newRouter() (*router_t){
	return &router_t{
		handlers: make(map[route_t]handleFunc),
		tries: make(map[string]*node),
	}
}

func (r *router_t) addRoute(route route_t, handler handleFunc) {
	root,ok := r.tries[route.method]
	if !ok {
		root = newNode("",false)
		r.tries[route.method] = root
	}
	root.insertPattern(parsePath(route.pattern),route.pattern,0)
	r.handlers[route] = handler
	log.Printf("Established Route %s - %s",route.method,route.pattern)
}

func (r *router_t) retrieve(method string, path string) (pattern string,params map[string]string) {
	pattern = ""
	params = make(map[string]string)
	trie, ok := r.tries[method]
	if !ok {
		return
	}
	names := parsePath(path)
	pattern = trie.getPattern(names,0)
	for i,v := range parsePath(pattern) {
		if len(v)>0 && v[0]==':' {
			params[v[1:]] = names[i]
		} else if len(v)>0 && v[0]=='*' {
			params[v[1:]] = strings.Join(names[i:],"/")
			return
		}
	}
	return
}


func (r *router_t) handle(c *Context) {
	stub, ok := r.handlers[route_t{
		method: c.method,
		pattern: c.pattern,
	}]
	log.Printf("Handle %s - %s",c.request.Host, c.request.URL)
	if ok {
		stub(c)
	} else {
		fmt.Fprintf(c.w,"404 Not Found : %s\n",c.request.URL)
	}
}