package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my lemonade stand")
	
}

func main() {
	fileServer := http.FileServer(http.Dir("static"))
	handler := http.StripPrefix("/static/", fileServer)
	http.HandleFunc("/", greet)
	http.Handle("/static/", handler)
	http.ListenAndServe(":8080", nil)
}
