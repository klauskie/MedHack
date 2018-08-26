package controllers

import (
	"log"
	"medhackmty/def"
	"medhackmty/medModelo"
	"medhackmty/sintModelo"
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

	lista, err := sintModelo.GetSintomas()
	if err != nil {
		log.Fatalln("Error, GetSintomas, ", err)
	}

	err = def.TPL.ExecuteTemplate(w, "sint.html", struct {
		Usr user
		Lst map[int]sintModelo.Sintoma
	}{
		dud,
		lista,
	})
	if err != nil {
		log.Fatalln(err)
	}
}

// GetSymptoms :
func GetSymptoms(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Redirect(w, r, "/", 303)
		return
	}

	r.ParseForm()
	sympSelected := r.Form["same"]
	enfermedadID := medModelo.GetEnf(sympSelected)
	medicamento, err := medModelo.GetMed(enfermedadID)
	if err != nil {
		log.Fatalln(err)
	}

	err = def.TPL.ExecuteTemplate(w, "resultado.html", medicamento)
	if err != nil {
		log.Fatalln(err)
	}

}
