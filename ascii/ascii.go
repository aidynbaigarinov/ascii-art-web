package ascii

import (
	"io/ioutil"
	"strings"
)

type Ascii struct {
	Input string
	Font  string
}

func CreateOutput(output, file []byte, word string, m int) []byte {
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
	return CreateOutput(output, file, word, m+1)
}

func AsciiOutput(a *Ascii) []byte {
	var (
		wordsArr []string
		output   []byte
		index    int
	)
	file, _ := ioutil.ReadFile("assets/banners" + a.Font + ".txt")
	wordsArr = strings.Split(a.Input, "\\n")
	for _, word := range wordsArr {
		output = CreateOutput(output, file, word, index)
	}
	return output
}
