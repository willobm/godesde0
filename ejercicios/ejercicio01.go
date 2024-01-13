package ejercicios

import (
	"strconv"
)

func ConviertoaEntero(strNumero string) (int, string) {

	Entero, err := strconv.Atoi(strNumero)
	if err != nil {
		return 0, "hubo el siguiente error:" + err.Error()
	} else {
		if Entero > 100 {
			return Entero, "Es mayor a 100"
		} else {
			return Entero, "Es menor a 100"
		}
	}

}
