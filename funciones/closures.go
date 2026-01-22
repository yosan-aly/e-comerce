package funciones

import "fmt"

func tabla(valor int) func() int {
	numero := valor
	secuencia := 0

	return func() int {
		secuencia++
		return numero * secuencia
	}
}

func LlamarClosure() {     //funcion que devuelve una funcion
	tabladel1 := 2
	MiTabla := tabla(tabladel1)
	for i := 1; i <= 10; i++ {
		fmt.Println(MiTabla())
	}
}

