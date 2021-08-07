package main

import "fmt"

var tabla [10]int
var matriz [5][7]int

func main() {
	/*tabla[0] = 1
	tabla[5] = 15
	fmt.Println(tabla)

	tabla2 := [10]int{5, 2, 3}
	fmt.Println(tabla2)

	for i := 1; i < len(tabla2); i++ {
		fmt.Println(tabla2[i])
	}*/
	slice := []int{2, 4, 6, 8, 10}
	porcion := slice[0:2]
	fmt.Println(porcion)
	variante()
}

func variante() {
	elementos := make([]int, 5, 20)
	fmt.Println(elementos)
	elementos = append(elementos, 5)
	fmt.Println(elementos)

	fmt.Printf("largo %d,capacidad %d", len(elementos), cap(elementos))

}
