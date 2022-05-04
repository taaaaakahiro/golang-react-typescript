package main

import (
	"fmt"
	"net/http"
	// "time"
	"html/template"
)

func main() {
	http.HandleFunc("/template", templateHandler)
	http.HandleFunc("/rest", restHandler)
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8080", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World!")
	fmt.Fprintln(w, "Hello World!")
}
// REST API
func restHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintln(w, "GET called!!")
	} else if r.Method == "POST" {
		fmt.Fprintln(w, "POST called!!")
	} else if r.Method == "PUT" {
		fmt.Fprintln(w, "PUT called!!")
	} else if r.Method == "DELETE" {
		fmt.Fprintln(w, "DELETE called!!")
	}
}

// Template Engine
func templateHandler(w http.ResponseWriter, r *http.Request){
	t, _ := template.ParseFiles("template/tmpl.html", "template/tmpl2.html")
	// t.Execute(w, time.Now())
	DaysOfWeek := []string{"日", "月", "火", "水", "木", "金", "土"}
	t.Execute(w, DaysOfWeek)
}
