package users

import (
	"fmt"
	"time"

	"github.com/yosan-aly/e-comerce/modelos"
)

func AltaUsuario() {
	u := new(modelos.User)

	u.AddUser(5, "Yohan", time.Now(), true)
	fmt.Println(u)
}
