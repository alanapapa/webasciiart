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

var text, fonts, Result string

func main() {
	http.HandleFunc("/", hello)
	fmt.Printf("Starting server for testing HTTP POST...\n http://127.0.0.1:1995/ \n")
	if err := http.ListenAndServe(":1995", nil); err != nil {
		log.Fatal(err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Error 404\nPage Not Found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "static")
	case "POST":
		fonts = r.FormValue("fonts")
		text = r.FormValue("text")

		for _, e := range text {
			if e < 32 || e > 126 {
				http.Error(w, "Error 400\nBad Request!", http.StatusNotFound)
				return
			}
		}

		file, err := os.Open("./" + fonts + ".txt")

		if err != nil {
			http.Error(w, "Internal server error - 500", http.StatusNotFound)
			return
		}

		fileVal := ScanFile(file)
		narg := strings.Split(text+" ", "\\n")
		for _, v := range narg {
			goResult(string(v), fileVal)
		}

		tmpl, _ := template.ParseFiles("static/index.html")
		tmpl.Execute(w, Result)
	}
}

func goResult(s string, fileVal []string) string {
	for i := 1; i <= 8; i++ {
		for _, arg := range s {
			indexBase := int(rune(arg)-32) * 9
			Result += string(fileVal[indexBase+i])
		}
		Result += "\n"
	}
	return Result
}

func ScanFile(font *os.File) []string {
	Result = ""
	var fileValue []string
	scanner := bufio.NewScanner(font)
	for scanner.Scan() {
		lines := scanner.Text()
		fileValue = append(fileValue, lines)
	}
	return fileValue
}
