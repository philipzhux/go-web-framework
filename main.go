package main
import(
	"gweb"
	"fmt"
)

func main() {
	gweb_instance := gweb.New()
	gweb_instance.GET("/test",func(w gweb.ResponseWriter,req *gweb.Request){
		fmt.Fprintf(w,"Test: %s",req.URL)
	})
	gweb_instance.GET("/",func(w gweb.ResponseWriter,req *gweb.Request){
		fmt.Fprintf(w,"Test at root: req.Method = %s",req.Method)
	})
	gweb_instance.Run("127.0.0.1:22222")
}
