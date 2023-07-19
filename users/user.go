package users

import (
	"fmt"
	"time"
	"godecero/modelos"
)

func AltaUsuario() {
	u := new(modelos.User)

	u.AddUser(10, "Pablo", "ivan@totaltools.cl", true, time.Now())

	fmt.Println(u)
}