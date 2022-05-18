package main

import (
	"fmt"
	"net/http"
	"html/template"
	"os"

	"github.com/joho/godotenv"
)

func Hello() {
	fmt.Println("Hello!")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World!!!!!!")
	fmt.Fprint(w, "Hello World!!!!!")
}

// REST API
func restHandler(w http.ResponseWriter, r *http.Request) {
	// Permit HTTP Request from localhost:3000(CORS)
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	// CASE1 ~switch文の場合~
	switch r.Method {
	case "GET":
		fmt.Fprint(w, "GET called!!")
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

func loadenv() {
	// .env 全体を読み込む
	err := godotenv.Load(".env")
	// 読み込み成功かどうかを確認
	if err != nil {
		fmt.Printf("読み込み失敗: %v", err)
	} 
	// .env内の SAMPLEを取得してmessageに代入
	message := os.Getenv("SAMPLE")

	fmt.Println(message)
}

func grammerHandler(w http.ResponseWriter, r *http.Request) {
// ポインタ
	p := 99 //　初期値
	pt := &p  //　ポインタ作成
	fmt.Println(*pt) // ポインタ呼び出し 99
	
// 構造体フィールド
	type vartex struct {
		sum int
		sumArray []int
	}
	var vt vartex
	fmt.Println(vt) // {0 []}
	fmt.Println(vt.sum) // 0
	fmt.Println(vt.sumArray) // []
	ptvt := &vt // 構造体型のポインタ作成
	for i := 0; i < 100; i++ {
		vt.sum += i
		vt.sumArray = append(vt.sumArray, i)
	}
	fmt.Println(*ptvt) // {499500 [0,1 ... 99]}
	// fmt.Println(*ptvt.sum) //NG
	fmt.Println(vt.sum) // 4050
	fmt.Println(vt.sumArray) // [0, 1 ... 99]
	fmt.Println(len(vt.sumArray)) // 100
	fmt.Println(vt.sumArray[5]) // 5
//for文
	//case1 ループ回数条件指定 
	var sum = 0
	for i := 0; i < 1000; i++ {
		sum += i	
	}
	//case2-1 //全て取り出し index不要な場合
	for i, _ := range vt.sumArray {
		fmt.Println(i) // 0 ~ 99
	}
	//case2-2 //全て取り出し indexを使用する場合
 	for _, v := range vt.sumArray {
		fmt.Println(v) // 0 ~ 99
	}
}