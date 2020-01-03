package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"time"
)

func main() {

	fmt.Println("It works")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		fmt.Println("Incoming request from " + r.RemoteAddr + " at ")
		fmt.Println(now)
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/golang", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Happy Gopher")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}
