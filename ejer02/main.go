package main

import (
	"fmt"
)

var numero int
var text string
var status bool

func main() {
	numero = 5
	//numero2 := 3

	//fmt.Println(numero2 + numero)
	//sumar(numero, numero2)

	fmt.Println("ingrese el numero de la rutina")
	fmt.Scanln(&numero)

	if numero == 1 {
		goto RUTINA1
	} else {
		goto RUTINA2
	}

RUTINA1:
	fmt.Println("esta es la rutina 1")

RUTINA2:
	fmt.Println("esta es la rutina 2")

}

func sumar(vnumero int, vnumero2 int) {
	fmt.Println(vnumero + vnumero2)
}
