package ejercicios

import (
	"strconv"
)

func Ejercicio01(texto string) (numero int, str string) {
	numero, err := strconv.Atoi(texto)
	if err != nil {
		return 0, "Error al convertir el texto a número"
	}

	if numero > 100 {
		return numero, "Es mayor a 100!"
	} else {
		return numero, "Es menor a 100"
	}
	/*
		switch numero {
		case numero > 100:
			fmt.Printf("Es mayor a 100")
		default:
			fmt.Printf("Es menor a 100")
		}*/
}
