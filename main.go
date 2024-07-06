package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my lemonade stand")	
}

func menuHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Below is our menu:")
	fmt.Fprintf(w, "Jug of lemonade: $50")
	fmt.Fprintf(w, "Glass of lemonade: $10")
	fmt.Fprintf(w, "Freshly-baked cup cakes: $5")
}

func main() {
	fileServer := http.FileServer(http.Dir("static"))
	handler := http.StripPrefix("/static/", fileServer)
	http.HandleFunc("/", homeHandler)
	http.Handle("/static/", handler)
	http.ListenAndServe(":8080", nil)
}
