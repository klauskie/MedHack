package medModelo

import (
	"fmt"
	"medhackmty/def"
)

type Medicamento struct {
	Id          int
	Nombre      string
	Descripcion string
}

func GetEnf(sintomas []string) int {
	if len(sintomas) == 1 {
		if sintomas[0] == "1" {
			return 5
		} else if sintomas[0] == "2" {
			return 6
		} else if sintomas[0] == "3" {
			return 7
		} else if sintomas[0] == "4" {
			return 8
		} else if sintomas[0] == "5" {
			return 9
		} else if sintomas[0] == "6" {
			return 10
		} else if sintomas[0] == "7" {
			return 11
		} else if sintomas[0] == "8" {
			return 12
		}
	} else {
		if sintomas[0] == "1" && sintomas[1] == "2" && sintomas[2] == "9" {
			return 13
		} else if sintomas[0] == "1" && sintomas[1] == "2" && sintomas[2] == "4" {
			return 14
		}
	}
	return 0
}

// GetMed : returns medicamento segun el id de la infermedad
func GetMed(idEnf int) (map[int]Medicamento, error) {
	fmt.Println("Dentro de getMed", idEnf)
	m := make(map[int]Medicamento)
	rows, err := def.DB.Query("SELECT medicamentos.id, medicamentos.nombre, medicamentos.descripcion FROM enf_med join medicamentos on enf_med.id_med = medicamentos.id where id_enf = ?", idEnf)
	if err != nil {
		fmt.Println("Error en hacer query...")
		return m, err
	}
	for rows.Next() {
		mdcmnt := Medicamento{}

		err := rows.Scan(&mdcmnt.Id, &mdcmnt.Nombre, &mdcmnt.Descripcion)
		if err != nil {
			return m, err
		}

		m[mdcmnt.Id] = mdcmnt
	}
	return m, nil
}
