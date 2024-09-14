package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hi, Welcome to Web Dev with Golang</h1>")
	fmt.Fprintf(w, "<h1>I love Golang</h1>")
	fmt.Fprintf(w, "<h1>I also love C++ & Rust</h1>")
	fmt.Fprintf(w, "<h1>I will eventually build my own website</h1>")

}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting the server on :3000")
	http.ListenAndServe("localhost:3000", nil)
}
