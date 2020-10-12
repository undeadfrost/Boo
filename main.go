package main

import (
	"Boo"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func main() {
	boo := Boo.New()
	boo.GET("/hello", Hello)
	boo.Run(":8080")
}
