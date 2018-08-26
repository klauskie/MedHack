package sintModelo

import (
	"medhackmty/def"
)

type Sintoma struct {
	Id     int
	Nombre string
}

func GetSintomas() (map[int]Sintoma, error) {
	m := make(map[int]Sintoma)
	rows, err := def.DB.Query("SELECT * FROM enfermedades")
	if err != nil {
		return m, err
	}
	for rows.Next() {
		sntm := Sintoma{}

		err := rows.Scan(&sntm.Id, &sntm.Nombre)
		if err != nil {
			return m, err
		}

		m[sntm.Id] = sntm
	}

	return m, nil
}
