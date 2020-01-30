package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
)

type Ascii struct {
	Input string
	Fs    string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	a := Ascii{}
	a.Input = r.FormValue("input")
	a.Fs = r.FormValue("fs")
	t.ExecuteTemplate(w, "index", string(asciiOutput(&a)))

}

func createOutput(output, file []byte, word string, m int) []byte {
	if m == 8 {
		return output
	}

	for j, n := 0, len(word); j < n; j++ {
		numOfNl := 0
		a := int(word[j]-32)*9 + 2 + m
		for i, l := 0, len(file); i < l; i++ {
			if file[i] == '\n' {
				numOfNl++
			} else if numOfNl == a-1 {
				output = append(output, file[i])
			} else if numOfNl == a {
				break
			}
		}
	}
	output = append(output, '\n')
	return createOutput(output, file, word, m+1)
}

func asciiOutput(a *Ascii) []byte {
	var (
		wordsArr []string
		output   []byte
		index    int
	)
	file, _ := ioutil.ReadFile("./assets/banners/" + a.Fs + ".txt")
	wordsArr = strings.Split(a.Input, "\\n")
	for _, word := range wordsArr {
		output = createOutput(output, file, word, index)
	}
	return output
}

func main() {
	fmt.Println("Listening on port: 8080")

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	http.HandleFunc("/", indexHandler)

	http.ListenAndServe(":8080", nil)
}
