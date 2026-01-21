package iteraciones

import "fmt"

func Iterar() {
	for i := 0; i < 100; i +=10 {
		fmt.Println(i)
	}

	for j := 100; j >= 0; j -=10 {
		fmt.Println(j)
	}

	for h := 0; h < 10; h++ {
		if h == 5 {
			continue // cuando llego a 5 volvio al for saltandose el 5. break para romper en 5
		}
		fmt.Println(h)
	}
}