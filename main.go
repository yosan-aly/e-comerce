package main

import (
	interfaces "github.com/yosan-aly/e-comerce/ejer_interfaces"
	"github.com/yosan-aly/e-comerce/modelos"
)


func main() {
	/*estado, texto := variables.CovertirATexto(123)

	fmt.Println(estado)
	fmt.Println(texto)
	if os := runtime.GOOS; os == "Linux." || os == "OS X." {
		fmt.Println("Sistema operativo Linux o MacOS")
	} else if os == "windows" {
		fmt.Println("Sistema operativo Windows")
	}*/

	/*switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Sistema operativo Linux")
	case "darwin":
		fmt.Println("Sistema operativo MacOS")
	default:
		fmt.Printf("%s\n", os)
	}*/

	/*numero, texto := ejercicios.Ejercicio01("Hola")

	fmt.Println(numero)
	fmt.Println(texto)

	teclado.IngresoNumeros()*/

	//fmt.Println(ejercicios.TablaNumerica())

	//files.SaveTable()

	//files.SumarTable()

	//files.ReadFile()

	//funciones.Calculos()

	//funciones.Exponencial(2)

	//arrellos_slices.Capacidad()

	//mapas.MostrarMapas()

	//users.AltaUsuario()
	
	yohan := new(modelos.Hombre)
	interfaces.HumanosPensando(yohan)

	sanly := new(modelos.Mujer)
	interfaces.HumanosPensando(sanly)
}
