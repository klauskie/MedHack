package controllers

import (
	"log"
	"medhackmty/def"
	"net/http"
)

// Home : Main page
func Home(w http.ResponseWriter, r *http.Request) {
	err := def.TPL.ExecuteTemplate(w, "home.html", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
