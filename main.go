package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Router struct{}

// ServeHTTP implements http.Handler.
func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	default:
		http.Error(w, "Page not found", http.StatusNotFound)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hi, Welcome to Web Dev with Golang</h1>")
	fmt.Fprintf(w, "<h1>I love Golang</h1>")
	fmt.Fprintf(w, "<h1>I also love C++ & Rust</h1>")
	fmt.Fprintf(w, "<h1>I will eventually build my own website</h1>")

}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>My email: <a href=\"mailto:abdullah.alkheshen@gmail.com\">abdullah.alkheshen@gmail.com</a>.</p>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>FAQ Page</h1><p>Here are the frequently asked questions.</p>")
}

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
