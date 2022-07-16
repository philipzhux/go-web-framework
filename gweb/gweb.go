package gweb
import (
	"log"
	"net/http"
)


type Engine struct{
	*Group
	router *router_t
	groups []*Group
}

type handleFunc func(*Context)

func New() (*Engine) {
	log.SetPrefix("[gweb] ")
	eng := &Engine{
		router:newRouter(),
		groups: make([]*Group,0),
		Group: &Group{},
	}
	eng.middleWares = make([]handleFunc, 0)
	eng.engine = eng
	return eng
}

/* implement the handler interface */
func (e *Engine) ServeHTTP(w http.ResponseWriter,req *http.Request) {
	e.router.handle(newContext(w,req,e))
}


/* entry point */
func (e *Engine) Run(addr string) (error) {
	return http.ListenAndServe(addr,e)
}
