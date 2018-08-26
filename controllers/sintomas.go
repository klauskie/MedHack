package controllers

import (
	"fmt"
	"log"
	"medhackmty/def"
	"net/http"
)

const TOT_SYMP = 4

type user struct {
	Age     string
	Sex     string
	SympMap map[string]bool
}

type diagnosis struct {
	D_Cabeza   bool
	D_Muscular bool
	Mareo      bool
	Vomito     bool
}

// GetInit : Get age and sex from user
func GetInit(w http.ResponseWriter, r *http.Request) {
	dud := user{}
	dud.SympMap = make(map[string]bool)

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

// GetSymptoms :
func GetSymptoms(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Redirect(w, r, "/", 303)
		return
	}

	r.ParseForm()
	symp_selected := r.Form["same"]
	medicamento := diagnosticar(symp_selected)

	err := def.TPL.ExecuteTemplate(w, "resultado.html", medicamento)
	if err != nil {
		log.Fatalln(err)
	}

}

// Diagnosticar
func diagnosticar(sintomas []string) string {
	fmt.Println("dentro de diagnosticar...")
	fmt.Println(sintomas)
	for _, val := range sintomas {
		fmt.Println(val)
		if val == "D_Cabeza" {
			return "Aspirina"
		} else if val == "D_Muscular" {
			return "Paracetamol"
		} else if val == "Mareo" {
			return "Algo para el mareo"
		} else {
			return "Algo para el vomito"
		}
	}
	return "nil"
}
