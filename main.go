package main

import "gweb"

func first[T any](a T, _ ...interface{}) T {
	return a
}
func main() {
	gweb_instance := gweb.New()
	gweb_instance.GET("/json/:test_param/b/*test_path",func(c *gweb.Context){
		err := c.SendJSON(gweb.StatusOK,gweb.AnyMap{
			"Method":c.GetMethod(),
			"test_param": first(c.GetParmas("test_param")),
			"test_path": first(c.GetParmas("test_path")),
			"Host": c.GetRequest().Host,
			"Agent": c.GetRequest().UserAgent(),
		})
		if err!= nil{
			c.SendString(gweb.StatusOK,"error: %v\n", err)
		}
	})
	gweb_instance.GET("/json/:test2/a/",func(c *gweb.Context){
		err := c.SendJSON(gweb.StatusOK,gweb.AnyMap{
			"Method":c.GetMethod(),
			"test2": first(c.GetParmas("test2")),
			"Host": c.GetRequest().Host,
			"Agent": c.GetRequest().UserAgent(),
		})
		if err!= nil{
			c.SendString(gweb.StatusOK,"error: %v\n", err)
		}
	})
	gweb_instance.POST("/post",func(c *gweb.Context){
		c.SendString(gweb.StatusOK,"String test: %s",c.GetRequestPost("value"))
	})
	gweb_instance.Run("127.0.0.1:22222")
}
