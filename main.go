package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {
	r := gee.New()

	r.GET("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "URL.Path = %q\n", r.URL.Path)
	})

	r.GET("/hello", func(rw http.ResponseWriter, r *http.Request) {
		for k, v := range r.Header {
			fmt.Fprintf(rw, "Header[%q] = %q\n", k, v)
		}
	})

	r.Run(":9999")
}
