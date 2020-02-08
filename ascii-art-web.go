package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

var text string
var fonts string
var Result string

func hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "static")
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		fonts = r.FormValue("fonts")
		text = r.FormValue("text")
		AsciiMain(text, fonts)
		tmpl, _ := template.ParseFiles("static/index.html")
		tmpl.Execute(w, Result)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {

	http.HandleFunc("/", hello)

	fmt.Printf("Starting server for testing HTTP POST...\n http://127.0.0.1:1995/ \n")
	if err := http.ListenAndServe(":1995", nil); err != nil {
		log.Fatal(err)
	}
}

func AsciiMain(text string, fonts string) {
	for _, e := range text {
		if e == '§' || e == '±' || e == '”' || e == '“' {
			text = "Lol Kek Cheburek"
		}
	}
	if fonts == "Thinkertoy" {
		file, _ := os.Open("thinkertoy.txt")
		fileVal := ScanFile(file)
		var test []string
		arg := text
		for _, v := range test {
			arg += " " + v
		}
		narg := strings.Split(arg, "\\n")
		for _, v := range narg {
			printLetter(string(v), fileVal)
		}
	} else if fonts == "Shadow" {
		file, _ := os.Open("shadow.txt")
		fileVal := ScanFile(file)
		var test []string
		arg := text
		for _, v := range test {
			arg += " " + v
		}
		narg := strings.Split(arg, "\\n")
		for _, v := range narg {
			printLetter(string(v), fileVal)
		}
	} else if fonts == "Standard" {
		file, _ := os.Open("standard.txt")
		fileVal := ScanFile(file)
		var test []string
		arg := text
		for _, v := range test {
			arg += " " + v
		}
		narg := strings.Split(arg, "\\n")
		for _, v := range narg {
			printLetter(string(v), fileVal)
		}
	} else {
		fmt.Println("Problem")
	}

	return

}

func printLetter(s string, fileVal []string) string {
	for i := 1; i <= 8; i++ {
		for _, arg := range s {
			indexBase := int(rune(arg)-32) * 9
			Result += string(fileVal[indexBase+i])
		}
		Result += "\n"
	}
	return Result
}

func ScanFile(arg *os.File) []string {
	Result = ""
	var fileValue []string
	scanner := bufio.NewScanner(arg)
	for scanner.Scan() {
		lines := scanner.Text()
		fileValue = append(fileValue, lines)
	}
	return fileValue
}
