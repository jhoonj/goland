package main

import (
	"fmt"
	"time"
)

type usuario struct {
	Id        int
	Nombre    string
	FechaAlta time.Time
	Status    bool
}

func main() {
	user := new(usuario)

	user.Id = 10
	user.Nombre = "jhon escobar"

	fmt.Println(user)

}
