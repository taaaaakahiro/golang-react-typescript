package main

import (
	"api/practice"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	// use func in handler.go
	Hello()
	// use packcage practice
	practice.Display()
	// .env
	loadenv()

	// chi
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	r.Use(middleware.Logger)
	r.Get("/welcome", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	r.Get("/", helloHandler)
	r.Get("/tmp", templateHandler)
	r.Get("/grm", GrammerHandler)
	// main
	r.Route("/param", func(r chi.Router) {
		// sub
		r.Route("/{ID}", func(r chi.Router) {
			r.Get("/", getParam)
			r.Post("/", postParam)     // POST /param/111
			r.Put("/", updateParam)    // PUT /param/111
			r.Delete("/", deleteParam) // DELETE /aparam/111
		})
	})
	http.ListenAndServe(":8080", r)

	// net/http
	// http.HandleFunc("/", helloHandler)
	// http.HandleFunc("/template", templateHandler)
	// http.HandleFunc("/rest", restHandler)
	// http.HandleFunc("/grammer", grammerHandler)
	// http.ListenAndServe(":8080", nil)

}

// handler with param
func getParam(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "ID")
	fmt.Println(ID)
	fmt.Fprintln(w, "GET /param/"+ID)
}
func postParam(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "ID")
	fmt.Println(ID)
	fmt.Fprintln(w, "POST /param/"+ID)
}
func updateParam(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "ID")
	fmt.Println(ID)
	fmt.Fprintln(w, "PUT /param/"+ID)
}
func deleteParam(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "ID")
	fmt.Println(ID)
	fmt.Fprintln(w, "DELETE /param/"+ID)
}
