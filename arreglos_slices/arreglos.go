package arrellos_slices

import "fmt"

var tabla [10]int = [10]int{10, 0, 59}
var matriz [3][3]int

func MuestroArreglos() {
	tabla[7] = 33
	tabla[2] = 15

	tabla2 := [10]string{"Yohan", "Aly"}

	fmt.Println(tabla)
	fmt.Println(tabla2)

	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}

	matriz[2][2] = 50
	fmt.Println(matriz)
}
