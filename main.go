package main

import (
	"fmt"
	"net/http"
	"time"
)

func Greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {

	http.HandleFunc("/", Greet)
	http.ListenAndServe("80", nil)
}

// func greet(w http.ResponseWriter, r *http.Request) {
//     fmt.Fprintf(w, "Hello World! %s", time.Now())
// }
