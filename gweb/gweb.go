package gweb
import (
	"log"
	"net/http"
)


type Engine struct{
	router *router_t
}

type handleFunc func(*Context)

func New() (*Engine) {
	log.SetPrefix("[gweb] ")
	return &Engine{newRouter()}
}

/* implement the handler interface */
func (e *Engine) ServeHTTP(w http.ResponseWriter,req *http.Request) {
	e.router.handle(newContext(w,req,e))
}


/* entry point */
func (e *Engine) Run(addr string) (error){
	return http.ListenAndServe(addr,e)
}

/* router setting interface exposed to users */
func (e *Engine) GET(pattern string, handler handleFunc) {
	e.router.addRoute(route_t{method: "GET",pattern: pattern},handler)
}

func (e *Engine) POST(pattern string, handler handleFunc) {
	e.router.addRoute(route_t{method: "POST",pattern: pattern},handler)
}