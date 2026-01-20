package ejercicios

import (
	"fmt"
	"strconv"
)

func Ejercicio01(texto string) (numero int, str string) {
	numero, _ = strconv.Atoi(texto)

	if numero > 100 {
		fmt.Println("Es mayor a 100")
	} else {
		fmt.Println("Es menor a 100")
	}
	/*
	switch numero {
	case numero > 100:
		fmt.Printf("Es mayor a 100")
	default:
		fmt.Printf("Es menor a 100")
	}*/
	return numero, texto
}
