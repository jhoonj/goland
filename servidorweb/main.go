package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/usuario", usuario)
	http.ListenAndServe(":7003", nil)
}

func home(w http.ResponseWriter, r *http.Request) { // para envio de datos a traves de un servidor siempre va a ser asi.
	http.ServeFile(w, r, "./index.html")
}

func usuario(w http.ResponseWriter, r *http.Request) { // para envio de datos a traves de un servidor siempre va a ser asi.
	http.ServeFile(w, r, "./usuario.html")
}
