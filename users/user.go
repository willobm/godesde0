package users

import (
	"fmt"
	"time"

	"github.com/willobm/godesde0/modelos"
)

func AltaUsuarios() {
	u := new(modelos.User) //en este momento se esta creando el objeto u de tipo user, no es una definición
	u.AddUser(23, "Pruebas", time.Now(), true)
	fmt.Println(u)
}
