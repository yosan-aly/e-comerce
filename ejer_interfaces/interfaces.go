package ejer_interfaces

import (
	"fmt"

	"github.com/yosan-aly/e-comerce/interfaces"
)

func HumanosPensando(hu interfaces.Humano) {
	hu.Pensar()
	fmt.Printf("Soy un/a %s, y estoy pensando \n", hu.Sexo())
}
