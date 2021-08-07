package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	go miNombreLento("jhon jairo")
	fmt.Println("asincrono")
	x := ""
	fmt.Scanln(&x)

}

func miNombreLento(nombre string) {

	archivo, _ := os.Create("./nuevoarchivo")
	letras := strings.Split(nombre, "")
	for _, letra := range letras {
		time.Sleep(2000 * time.Millisecond)
		fmt.Fprintln(archivo, letra)
	}

	archivo.Close()
}
