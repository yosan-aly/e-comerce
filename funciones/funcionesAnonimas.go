package funciones

import "fmt"

func Calculos() {

	calculo := func(numero1 int, numero2 int) int {
		return numero1 + numero2
	}

	fmt.Println(calculo(10, 5))

	calculo = func(numero1 int, numero2 int) int {  //funcon anonima
		return numero1*numero2
	}

	fmt.Println(calculo(1,2))
}
