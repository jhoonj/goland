package main

import "fmt"

func main() {
	fmt.Println("entro")
	exponente(5, 3)

}

func exponente(numero int, base int) {
	if base > 1 {
		fmt.Println("numero", numero)
		exponente(numero*numero, base-1)

	} else {
		fmt.Println("resultado", numero)
		return
	}
}
