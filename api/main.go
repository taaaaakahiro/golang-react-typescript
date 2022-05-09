package main

import (
	"fmt"
	"net/http"
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
	// Permit HTTP Request from localhost:3000(CORS)
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	// CASE1 ~switch文の場合~
	switch r.Method {
	case "GET":
		fmt.Fprintln(w, "GET called!!")
	case "POST":
		fmt.Fprintln(w, "POST called!!")
	case "PUT":
		fmt.Fprintln(w, "UPDATE called!!")
	case "DELETE":
		fmt.Fprintln(w, "DELETE called!!")
	default:
		w.WriteHeader(405)
	}
	// CASE2 ~if文の場合~
	// if r.Method == "GET" {
	// 	fmt.Fprintln(w, "GET called!!")
	// } else if r.Method == "POST" {
	// 	fmt.Fprintln(w, "POST called!!")
	// } else if r.Method == "PUT" {
	// 	fmt.Fprintln(w, "PUT called!!")
	// } else if r.Method == "DELETE" {
	// 	fmt.Fprintln(w, "DELETE called!!")
	// }
}

// Template Engine
func templateHandler(w http.ResponseWriter, r *http.Request){
	t, _ := template.ParseFiles("template/tmpl.html", "template/tmpl2.html")
	DaysOfWeek := []string{"日", "月", "火", "水", "木", "金", "土"}
	t.Execute(w, DaysOfWeek)
}

func testPractice(x, y int) int {
	z := x + y
	return z
}
