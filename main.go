package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	// Initialize
	r := mux.NewRouter()
	fmt.Println("Initialize go server: http://localhost:8080")

	// Set routes with handlers
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/hoge", hogeHandler)
	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	// Start Server
	r.Use(loggingMiddleware)
	http.ListenAndServe(":8080", r)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, from root path\n"))
}

func hogeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla !!!\n"))
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("NOT FOUND\n"))
}

// define loggin middleware
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.RequestURI)

		next.ServeHTTP(w, r)
	})
}
