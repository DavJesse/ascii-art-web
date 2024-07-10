package Web

import (
	"net/http"
)

func FormHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/submitForm.html")
}

func SubmitFormHandler(w http.ResponseWriter, r http.Request) (string, string) {
	var bnStyle, inputStr string
	if r.Method == http.MethodPost {
		bnStyle = r.FormValue("style")
		inputStr = r.FormValue("inputStr")
	} else {
		http.Error(w, "Invalid Request Method!", http.StatusMethodNotAllowed)
	}
	return bnStyle, inputStr
}
