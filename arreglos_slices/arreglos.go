package arreglos_slices

import "fmt"

var tabla [10]int = [10]int{6, 5, 3, 2, 3}
var matriz [20][30]int

func MuestroArreglo() {
	tabla[7] = 33
	tabla[2] = 54
	fmt.Println(tabla)
	tabla2 := [10]string{"juan", "pedro", "lucas"}
	fmt.Println(tabla2)
	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}
	matriz[6][28] = 444
	fmt.Println(matriz)
}
