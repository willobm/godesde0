package main

import (
	"fmt"
	"runtime"

	"github.com/willobm/godesde0/variables"
)

func main() {
	estado, texto := variables.ConveritoaTexto(100)
	if estado {
		fmt.Println(texto)
	} else {
		fmt.Println("No pudo convertir a texto")
	}

	if os := runtime.GOOS; os == "linux" || os == "OS X." {
		fmt.Println("No es un Windows, es: ", os)
	} else {
		fmt.Println("Es un Windows: ", os)
	}

	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("Esto es un windows")
	case "darwin":
		fmt.Println("Esto es un darwin")
	default:
		fmt.Printf("%s \n", os)
	}
}
