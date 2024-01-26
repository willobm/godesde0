package users

import (
	"fmt"
	"time"

	"github.com/willobm/godesde0/modelos"
)

func AltaUsuarios() {
	u := new(modelos.User)
	u.AddUser(23, "Pruebas", time.Now(), true)
	fmt.Println(u)
}
