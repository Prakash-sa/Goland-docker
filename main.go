package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", HandleFunction)
	http.HandleFunc("/hi", HiFunction)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func HandleFunction(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func HiFunction(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi")
}
