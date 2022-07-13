package gweb
import (
	"log"
	"net/http"
	"fmt"
)
type ResponseWriter http.ResponseWriter
type Request http.Request
type HandleFunc func(ResponseWriter,*Request)
type route_t struct{
	Method string;
	Pattern string
}

type Engine struct{
	router map[route_t]HandleFunc
}

func New() (*Engine) {
	log.SetPrefix("[gweb]")
	return &Engine{make(map[route_t]HandleFunc)}
}

/* implement the handler interface */
func (e *Engine) ServeHTTP(w http.ResponseWriter,req *http.Request) {
	stub, ok := e.router[route_t{
		Method: req.Method,
		Pattern: req.URL.Path,
	}]
	log.Printf("Handle %s - %s",req.Host, req.URL)
	if ok {
		stub(ResponseWriter(w),(*Request)(req))
	} else {
		fmt.Fprintf(w,"404 Not Found : %s\n",req.URL)
	}
}


func (e *Engine) addRoute(route route_t, handler HandleFunc) {
	e.router[route] = handler
	log.Printf("Established Route %s - %s",route.Method,route.Pattern)
}

func (e *Engine) GET(route string, handler HandleFunc) {
	e.addRoute(route_t{Pattern: route, Method: "GET"},handler)
}

func (e *Engine) POST(route string, handler HandleFunc) {
	e.addRoute(route_t{Pattern: route, Method: "POST"},handler)
}

/* entry point */
func (e *Engine) Run(addr string) (error){
	return http.ListenAndServe(addr,e)
}
