package main

import (
	"log"
)

func main() {
	//_, err := os.Open("archivosx")
	//defer fmt.Println("se ejecuto ")    //defer es una finali

	//fmt.Println(err)

	ejemploPanic()
}

func ejemploPanic() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("ocurrio un error de panic", reco)
		}
	}()

	a := 1
	if a == 1 { // el panic me aborta el mensaje
		panic("se encontro el valor de 1")
	}
}
