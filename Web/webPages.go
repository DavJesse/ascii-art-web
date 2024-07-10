package Web

import "net/http"

func FormHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/submitForm.html")
}
