package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TablaNumerica() {
	var numero int
	var err error
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Ingrese un número:")
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("El dato ingresado es incorrecto:" + err.Error())
				continue
			} else {
				break
			}
		}
	}

	for i := 0; i <= 10; i++ {
		fmt.Printf("%d X %d = %d \n", numero, i, numero*i)
	}
}
