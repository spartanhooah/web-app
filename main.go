package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", homePage)

	log.Println("Starting web server on port 8080")
	http.ListenAndServe(":8080", nil)
}

func homePage(response http.ResponseWriter, request *http.Request) {
	renderTemplate(response, "index.html")
}

func renderTemplate(w http.ResponseWriter, page string) {
	t, err := template.ParseFiles(page)

	if err != nil {
		log.Println(err)
		return
	}

	err = t.Execute(w, nil)
	
	if err != nil {
		log.Println(err)
		return
	}
}