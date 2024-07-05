package main

import (
	"fmt"
	"net/http"
	//"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my first web server")
	//fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {
	fileServer := http.FileServer(http.Dir("static"))
	handler := http.StripPrefix("/static/", fileServer)
	http.HandleFunc("/", greet)
	http.Handle("/static/", handler)
	http.ListenAndServe(":8080", nil)
}
