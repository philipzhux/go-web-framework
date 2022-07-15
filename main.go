package main

import "gweb"

func main() {
	gweb_instance := gweb.New()
	gweb_instance.GET("/string",func(c *gweb.Context){
		c.SendString(gweb.StatusOK,"String test: %s",c.Request.URL)
	})
	gweb_instance.GET("/json",func(c *gweb.Context){
		err := c.SendJSON(gweb.StatusOK,gweb.AnyMap{
			"Method":c.Request.Method,
			3: c.Request.URL,
			"Host": c.Request.Host,
			"Agent": c.Request.UserAgent(),
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
