package arrellos_slices

import "fmt"


var tabla3 []int = []int{1,2,3}
var arreglo [10]int = [10]int{5,4,3,2,1,0,9,8,7,6}

func MuestroSlices() {
	fmt.Println(tabla3)

	porcion := arreglo[2:] //Slice desde la posicion 2 hasta el final
	porcion2 := arreglo[:7]
	porcion3 := arreglo[2:4]

	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}

func Capacidad() {
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))

	nums := make([]int, 0, 0)
	for i:= 0; i < 100; i++ {
		nums = append(nums, i)
	}

	fmt.Printf("Largo %d, Capacidad %d", len(nums), cap(nums))
}