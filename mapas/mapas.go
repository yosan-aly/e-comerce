package mapas

import "fmt"

func MostrarMapas() {
	paises := make(map[string]string)

	paises["Mexico"] = "D.F"
	paises["Colombia"] = "Bogota"
	fmt.Println(paises)
	fmt.Println(paises["Colombia"])

	campeonato := map[string]int{
		"Colombia": 10,
		"Mexico":   8,
		"Brasil":   9,
	}
	fmt.Println(campeonato)

	for equipo, puntaje := range campeonato {
		fmt.Printf("Equipo %s, tiene un puntaje de %d \n", equipo, puntaje)
	}

	delete(campeonato, "Mexico")
	fmt.Println(campeonato)

	puntajeConsulta, existe := campeonato["Colombia"]
	fmt.Printf("Puntaje consultado: %d, y el equipo existe = %t", puntajeConsulta, existe)


}