package main

import "fmt"

type serVivo interface {
	estaVivo() bool
}

type humano interface {
	respirar()
	pensar()
	comer()
	sexo() string
	estaVivo() bool
}

type animal interface {
	respirar()
	comer()
	esCarnivoro() bool
	estaVivo() bool
}

type vegetal interface {
	clasificacionVegetal() string
	estaVivo() bool
}

/* humano */

type hombre struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	esHombre   bool
	vivo       bool
}

type mujer struct {
	hombre
}

func (h *hombre) respirar() { h.respirando = true }
func (h *hombre) comer()    { h.comiendo = true }
func (h *hombre) pensar()   { h.pensando = true }
func (h *hombre) sexo() string {
	if h.esHombre == true {
		return "hombre"
	} else {
		return "mujer"
	}
}
func (h *hombre) estaVivo() bool { return h.vivo }

func HumanosRespirando(hu humano) {
	hu.respirar()
	fmt.Printf("soy un/a %s y ya estoy respirando \n", hu.sexo())

}

func estoyVivo(v serVivo) bool {
	return v.estaVivo()
}

func main() {
	Pedro := new(hombre)
	Pedro.vivo = true
	Pedro.esHombre = true
	HumanosRespirando(Pedro)
	fmt.Printf("estoy vivo %t pedro", estoyVivo(Pedro))

	Maria := new(mujer)
	HumanosRespirando(Maria)
	fmt.Printf("estoy vivo %t maria", estoyVivo(Maria))

}
