package main

import "fmt"

/*
Middlewares en GO son intercepterores qeu permiter ejecutar instrucciones comunes a varias funciones que reciben y devuelven los mismos tipos de variables
*/

var result int

func main() {
	fmt.Println("inicio")
	result = operacionesMidd(sumar)(2, 3)
	fmt.Println(result)
}

func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}

func operacionesMidd(f func(int, int) int) func(int, int) int {
	return func(i1, i2 int) int {
		fmt.Println("inicio de operacion")
		return f(i1, i2)
	}
}
