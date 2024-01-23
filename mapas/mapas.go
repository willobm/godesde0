package mapas

import "fmt"

func MostrarMapas() {
	paises := make(map[string]string)
	paises["Ecuador"] = "Quito"
	paises["Peru"] = "Lima"

	// fmt.Println(paises)
	// fmt.Println(paises["Ecuador"])

	campeonato := map[string]int{
		"LDU":         45,
		"Emelec":      42,
		"Nacional":    40,
		"Barcelona":   39,
		"U. Catolica": 39,
		"Aucas":       20,
	}
	fmt.Println(campeonato)

	// for equipo, puntaje := range campeonato {
	// 	fmt.Printf("El %s, tiene un puntaje de %d \n", equipo, puntaje)
	// }

	delete(campeonato, "Barcelona")

	fmt.Println(campeonato)

	// puntaje, existe := campeonato["Barcelona"]

	// fmt.Printf("El puntaje capturado es: %d, y el equipo existe=%t \n", puntaje, existe)
	puntaje, existe := campeonato["Aucas"]

	fmt.Printf("El puntaje capturado es: %d, y el equipo existe=%t \n", puntaje, existe)
}
