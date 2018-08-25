package controllers

import (
	"log"
	"medhackmty/def"
	"net/http"
)

type user struct {
	Age string
	Sex string
}

// GetInit : Get age and sex from user
func GetInit(w http.ResponseWriter, r *http.Request) {
	dud := user{}

	if r.Method != "POST" {
		http.Redirect(w, r, "/", 303)
		return
	}

	dud.Age = r.FormValue("age")
	dud.Sex = r.FormValue("gender")

	err := def.TPL.ExecuteTemplate(w, "sint.html", dud)
	if err != nil {
		log.Fatalln(err)
	}
}
