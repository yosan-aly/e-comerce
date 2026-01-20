package main

import (
	"fmt"

	"github.com/yosan-aly/e-comerce/variables"
)

func main() {
	estado, texto :=variables.CovertirATexto(123)

	fmt.Println(estado)
	fmt.Println(texto)
}