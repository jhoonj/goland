package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	leoArchivo()
	leoArchivo2()
	graboArchivo()
}

func leoArchivo() {
	archivo, err := ioutil.ReadFile("./archivo.txt") // no deja maniputar el archivo, no debo abrir y cerrrar el archivo
	if err != nil {
		fmt.Println("hubo un error")
	} else {
		fmt.Println(string(archivo))
	}
}

func leoArchivo2() {
	archivo, err := os.Open("./archivo.txt") //deja manipular el archivo si debo abrir y cerrar el archivo
	if err != nil {
		fmt.Println("hubo un error")
	} else {
		scanner := bufio.NewScanner(archivo)
		i := 0
		for scanner.Scan() {
			i++
			registro := scanner.Text()
			fmt.Printf("linea %d %s \n", i, registro)
		}
	}
	archivo.Close()
}

func graboArchivo() {
	archivo, err := os.Create("./archivoNuevo.txt") //deja manipular el archivo si debo abrir y cerrar el archivo
	if err != nil {
		fmt.Println("hubo un error")
		return
	} else {
		fmt.Fprintln(archivo, "esta es una linea nueva")
		archivo.Close()
	}
}
