package defer_panic

import (
	"fmt"
	"log"
)

func VemosDefer() {
	fmt.Println("Primer mensaje")
	defer fmt.Println("Último mensaje")
	fmt.Println("Segundo mensaje")
}

func EjemploPani() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("Ocurrio un error que genero panic \n %v", reco)
		}
	}()
	a := 1
	if a == 1 {
		panic("Encontre el valor de 1")
	}
}
