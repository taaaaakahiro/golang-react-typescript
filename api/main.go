package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// use func in handler.go
	Hello()
	// .env
	loadenv()
	// chi
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("welcome"))
	// })
	r.Get("/", helloHandler)
	r.Get("/tmp", templateHandler)
	r.Get("/grm", grammerHandler)
	http.ListenAndServe(":8080", r)

	// net/http
	// http.HandleFunc("/", helloHandler)
	// http.HandleFunc("/template", templateHandler)
	// http.HandleFunc("/rest", restHandler)
	// http.HandleFunc("/grammer", grammerHandler)
	// http.ListenAndServe(":8080", nil)




}