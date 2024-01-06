package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *httpRequest) {
	fmt.Fprintf(w, "<h1>Welcome to Web Dev with Go</h1>")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting the server on :3000")
	http.ListenAndServe(":3000", nil)
}
