package main

import "fmt"

var Calculo func(int, int) int = func(i1, i2 int) int {
	return i1 + i2
}

func main() { // las funciones anonimas no tienen nombres -- es como una especie de sobrecarga, pero no me deja modificar los parametros de entrada
	fmt.Println("sumo 5 +7 ", Calculo(5, 7))

	Calculo = func(i1, i2 int) int {
		return i1 - i2
	}

	fmt.Println("resto 5 - 7 ", Calculo(5, 7))

	Calculo = func(i1, i2 int) int {
		return i1 * i2
	}

	fmt.Println("multiplico 5 *7 ", Calculo(5, 7))

	operaciones()

	tabladel := 2
	MiTabla := Tabla(tabladel)
	for i := 1; i < 11; i++ {
		fmt.Println(MiTabla())
	}

}

func operaciones() {
	resultado := func() int {
		var a int = 23
		var b int = 13
		return a + b
	}
	fmt.Println(resultado())
}

func Tabla(valor int) func() int { //CLOSURES
	numero := valor
	secuencia := 0
	return func() int {
		secuencia = secuencia + 1
		return numero * secuencia
	}
}
