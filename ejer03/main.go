package main

import "fmt"

func main() {
	fmt.Println("compila")

	fmt.Println(uno(3))

	numero, estado := dos(3)
	println(numero)
	println(estado)

}

func uno(numero int) int {
	return numero * 2
}

func dos(numero int) (int, bool) {
	return numero * 2, false

}
