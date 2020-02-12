package main

import (
	"html/template"
	"log"
	"net/http"

	"./ascii"
)

type ExecOutput struct {
	In  string
	Out string
}

func ValidAscii(s string) bool {
	for _, i := range []byte(s) {
		if i > 127 {
			return false
		}
	}
	return true
}

func internalServerError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	t, _ := template.ParseFiles("internalerror.html")
	err := t.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		switch r.Method {
		case "GET":
			t, err := template.ParseFiles("index.html")
			if err != nil {
				internalServerError(w, r)
			}
			t.Execute(w, nil)
		case "POST":
			r.ParseForm()
			if !ValidAscii(r.Form.Get("input")) {
				w.WriteHeader(http.StatusBadRequest)
				t, err := template.ParseFiles("badrequest.html")
				if err != nil {
					internalServerError(w, r)
				}
				t.Execute(w, nil)
			} else {
				output, status := ascii.AsciiOutput(r.Form["input"][0], r.Form["font"][0])
				log.Printf("method: %v / font: %v / input: %v / statuscode: %v\n", r.Method, r.Form["font"][0], r.Form["input"][0], status)
				if status > 200 {
					internalServerError(w, r)
				} else {
					ex := ExecOutput{
						In:  r.Form["input"][0],
						Out: output,
					}
					t, err := template.ParseFiles("index.html")
					if err != nil {
						internalServerError(w, r)
						return
					}
					// fmt.Println(ex.Out)
					t.Execute(w, ex)
				}
			}
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		t, err := template.ParseFiles("notfound.html")
		if err != nil {
			internalServerError(w, r)
			return
		}
		t.Execute(w, nil)
	}
}

func main() {
	log.Println("server is starting...")
	http.HandleFunc("/", Handler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
