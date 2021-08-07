package main

import (
	"fmt"
)

var n1, n2, n3 int

func main() {

	//n1, n2, n3 := tresRetornos(1)

	calcula(6, 7, 8)

	/*	fmt.Println("listo", dos(6))
		fmt.Println(n1)
		fmt.Println(n2)
		fmt.Println(n3)*/
}

func tresRetornos(numero int) (int, int, int) {
	return numero, numero * numero, numero * numero * numero
}

func dos(numero int) int {
	return numero * 2
}

func calcula(numero ...int) { // de esta manera le digo que va a recibir parametros enteros pero no se cuantaos

	total := 0
	for posicion, valor := range numero { // range me devuele dos paratros siempre que en go no quiera usar valor de los que me retorna uso _
		fmt.Println("posicion ", posicion)
		fmt.Println("valor ", valor)
		total = total + valor
	}
	fmt.Println("total ", total)

}
