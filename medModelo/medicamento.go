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
