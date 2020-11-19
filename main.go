package main

import (
	"Boo"
	"fmt"
)

func Hello(c *Boo.Context) {
	name := c.Query("name")
	fmt.Println(name)
	user := map[string]string{
		"name": name,
		"age":  "27",
		"爱好":   "哭鼻子",
	}
	c.Json(200, user)
}

func main() {
	boo := Boo.New()
	boo.GET("/", Hello)
	boo.GET("/hello", Hello)
	boo.GET("/hello/:name", Hello)
	boo.GET("/hello/:suck", Hello)
	boo.GET("/hello2", Hello)
	boo.GET("/hello3/:name", Hello)
	boo.Run(":8080")
}
