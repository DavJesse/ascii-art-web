package main

import (
	"log"
	"net/http"
	"web/Web"
)

func main() {
	//	http.HandleFunc("/", Web.FormHandler)
	http.HandleFunc("/", Web.SubmitFormHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Fatal(http.ListenAndServe(":8000", nil))

}
