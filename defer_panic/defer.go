package defer_panic

import (
	"fmt"
	"log"
)

func VemosDefer() {
	fmt.Println("Inicio de la función primer mensaje")
	defer fmt.Println("Esto se ejecuta al final de la función defer")

	fmt.Println("tercer mensaje")
}

func EjemploPaninc() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("Ocurrio un error que genero panic \n %v", reco)
		}
	}()
	a := 1
	if a == 1 {
		panic("Se encontro el valor a")
	}
}
