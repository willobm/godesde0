package funciones

import "fmt"

func Calculos() {
	var numero3 int = 10
	var numero4 int = 54
	calculo := func(numero1 int, numero2 int) int {
		return numero1 + numero2 + numero3
	}

	fmt.Println(calculo(10, 44))

	calculo = func(numero1 int, numero2 int) int {
		return numero1 * numero2 * numero4
	}

	fmt.Println(calculo(3, 7))
}
