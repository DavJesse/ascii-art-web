package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my lemonade stand")
}

func menuHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Below is our menu:\n")
	fmt.Fprintf(w, "Jug of lemonade: $50\n")
	fmt.Fprintf(w, "Glass of lemonade: $10\n")
	fmt.Fprintf(w, "Freshly-baked cup cakes: $5\n")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About Us\n")
	fmt.Fprintf(w, "The Lemonade Inc was started in 2008 by David Odhiambo out of his sheer passion for lemonade. Over the years, the Lemonade Inc has contributed massively to the bevarage industry. It revolutionized how we approach and consume lemonade, introducing the world to the twirling straw.\n")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contacts\n")
	fmt.Fprintf(w, "Contact us at:\n")
	fmt.Fprintf(w, "Email: enquiries@lemonade.com\n")
	fmt.Fprintf(w, "Call: 0712 345 689\n")
}

func orderHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/order.html")
}

func submitOrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		quantity := r.FormValue("quantity")
		fmt.Fprintf(w, "Thank you %s! Your order for %s lemanades has been recieved.\n", name, quantity)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func main() {
	fileServer := http.FileServer(http.Dir("static"))
	handler := http.StripPrefix("/static/", fileServer)
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/menu", menuHandler)
	http.HandleFunc("/order", orderHandler)
	http.HandleFunc("/submit-order", submitOrderHandler)
	http.HandleFunc("/contacts", contactHandler)
	http.HandleFunc("/about", aboutHandler)
	http.Handle("/static/", handler)
	http.ListenAndServe(":8080", nil)
}
