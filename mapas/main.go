package main

import "fmt"

func main() {
	fmt.Println("mapas")

	paises := make(map[string]string) // CREACION DE MAPA
	paises["Mexico"] = "D.F."
	paises["ARGENTINA"] = "DEMO"

	fmt.Println(paises)

	campeonato := map[string]int{
		"barcelaona": 1,
		"chivas":     36,
		"medellin":   56}

	fmt.Println(campeonato)

	campeonato["river"] = 35

	fmt.Println(campeonato)

	delete(campeonato, "chivas")

	fmt.Println(campeonato)

	for equipo, puntage := range campeonato { // recorrer el mapa
		fmt.Printf("el equipo %s, tiene un puntaje %d \n", equipo, puntage)
	}

	puntaje, existe := campeonato["barcelaona"]
	fmt.Printf("tiene un puntaje %d y el equipo existe ? existe %t \n", puntaje, existe)

}
