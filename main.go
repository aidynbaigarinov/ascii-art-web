package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"./ascii"
)

func ValidLength(s string) bool {
	if len(s) == 0 {
		return false
	}
	return true
}

func ValidAscii(s string) bool {

}

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		r.ParseForm()
		t, _ := template.ParseFiles("index.html")
		err := t.Execute(w, nil)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("method: %v / font: %v / input: %v\n", r.Method, r.Form["font"], r.Form["input"])
		if !ValidLength(r.Form.Get("input")) {
			fmt.Println("type something")
		} else {
			a := &ascii.Ascii{
				Font:  r.Form["font"][0],
				Input: r.Form["input"][0],
			}
			fmt.Println(a)
		}

	} else {
		http.NotFound(w, r)
		return
	}
}

func main() {
	log.Println("server is starting...")
	http.HandleFunc("/", Handler)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal(err)
	}
}
