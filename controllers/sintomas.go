package controllers

import (
	"fmt"
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
	medicamento, err := diagnosticar(sympSelected)
	if err != nil {
		log.Fatalln(err)
	}

	err = def.TPL.ExecuteTemplate(w, "resultado.html", medicamento)
	if err != nil {
		log.Fatalln(err)
	}

}

// Diagnosticar
func diagnosticar(sintomas []string) (map[int]medModelo.Medicamento, error) {
	fmt.Println("dentro de diagnosticar...")
	fmt.Println(sintomas)
	if len(sintomas) < 2 {
		if sintomas[0] == "5" {
			return medModelo.GetMed(5)
		} else if sintomas[0] == "6" {
			return medModelo.GetMed(6)
		} else if sintomas[0] == "7" {
			return medModelo.GetMed(7)
		} else if sintomas[0] == "8" {
			return medModelo.GetMed(8)
		} else if sintomas[0] == "9" {
			return medModelo.GetMed(9)
		} else if sintomas[0] == "10" {
			return medModelo.GetMed(10)
		} else if sintomas[0] == "11" {
			return medModelo.GetMed(11)
		} else {
			return medModelo.GetMed(12)
		}
	}

	return nil, nil
}
