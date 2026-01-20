package variables

import "fmt"

func MuestroEnteros() {
	var intComun int
	intde32 := int32(10)
	intde64 := int64(100)

	fmt.Println("Valor int comun:", intComun)
	fmt.Println("Valor int32 convertido a int:", int(intde32))
	fmt.Println("Valor int64 convertido a int:", int(intde64))
}
