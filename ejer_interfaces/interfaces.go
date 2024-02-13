package ejer_interfaces

import (
	"fmt"

	"github.com/willobm/godesde0/interfaces"
)

func HumanosRespirando(hu interfaces.Humano) {
	fmt.Printf("Soy un/a %s, estoy respirando porque estoy vivo?:%t \n", hu.Sexo(), hu.Estavivo())
}
