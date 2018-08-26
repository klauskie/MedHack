package enfModelo

import (
	"medhackmty/def"
	"medhackmty/sintModelo"
	"strconv"
)

// Enfermedad : ___
type Enfermedad struct {
	Id       int
	Nombre   string
	Sintomas []sintModelo.Sintoma
}

// AuxPossible : ___
func AuxPossible(cadena []string) ([]int, error) {
	wrd, _ := strconv.Atoi(cadena[0])
	lista, err := PossibleEnf1(wrd)
	if err != nil {
		return nil, err
	}

	for i := range cadena {
		for j, val := range lista {
			wrd, err = strconv.Atoi(cadena[i+1])
			state, _ := Compare(val, wrd)
			if !state {
				lista = append(lista[:j], lista[j+1:]...)
			}
		}

	}
	return lista, nil
}

// PossibleEnf1 : ___
func PossibleEnf1(idSint int) ([]int, error) {
	lista := []int{}

	rows, err := def.DB.Query("select enfermedades.id from enfermedades join enf_sint on enfermedades.id = enf_sint.id_enf where enf_sint.id_sint = ?", idSint)
	if err != nil {
		return lista, err
	}
	for rows.Next() {
		enfer := Enfermedad{}

		err := rows.Scan(&enfer.Id)
		if err != nil {
			return lista, err
		}

		lista = append(lista, enfer.Id)
	}
	return lista, nil
}

// PossibleEnf2 : ___
func PossibleEnf2(idEnf int, idSint int) ([]int, error) {
	lista := []int{}

	rows, err := def.DB.Query("select enfermedades.id from enfermedades join enf_sint on enfermedades.id = enf_sint.id_enf where enf_sint.id_enf = ? and enf_sint.id_sint = ?", idEnf, idSint)
	if err != nil {
		return lista, err
	}
	for rows.Next() {
		enfer := Enfermedad{}

		err := rows.Scan(&enfer.Id)
		if err != nil {
			return lista, err
		}

		lista = append(lista, enfer.Id)
	}
	return lista, nil
}

// Compare : ___
func Compare(idEnf int, idSint int) (bool, error) {

	lista, err := PossibleEnf1(idSint)
	if err != nil {
		return false, err
	}
	for _, val := range lista {
		if val == idEnf {
			return true, nil
		}
	}

	return false, nil

}
