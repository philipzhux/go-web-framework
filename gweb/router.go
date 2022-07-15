package gweb
import (
	"log"
	"fmt"
)

type route_t struct{
	Method string;
	Pattern string
}

type router_t struct {
	handlers map[route_t]handleFunc
}

func newRouter() (*router_t){
	return &router_t{handlers: make(map[route_t]handleFunc)}
}

func (r *router_t) addRoute(route route_t, handler handleFunc) {
	r.handlers[route] = handler
	log.Printf("Established Route %s - %s",route.Method,route.Pattern)
}


func (r *router_t) handle(c *Context) {
	stub, ok := r.handlers[route_t{
		Method: c.Request.Method,
		Pattern: c.Request.URL.Path,
	}]
	log.Printf("Handle %s - %s",c.Request.Host, c.Request.URL)
	if ok {
		stub(c)
	} else {
		fmt.Fprintf(c.W,"404 Not Found : %s\n",c.Request.URL)
	}
}