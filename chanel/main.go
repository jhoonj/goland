package main

import (
	"fmt"
	"time"
)

func main() {
	canal1 := make(chan time.Duration)
	go bucle(canal1)
	fmt.Println("hasta aca llego")

	msg := <-canal1
	fmt.Println(msg)
}

func bucle(can chan time.Duration) {

	inicio := time.Now()
	for i := 0; i < 1000000000000; i++ {

	}
	final := time.Now()

	can <- final.Sub(inicio) // me da la duracion entre el fina y el fin y lo asigna al canal

}
