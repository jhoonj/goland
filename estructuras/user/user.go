package user

import "time"

type Usuario struct {
	Id        int
	Nombre    string
	FechaAlta time.Time
	Status    bool
}

// this significa que que hace referencia al objeto donde esta parado
func (this *Usuario) AltaUsuario(id int, nombre string, fechaAlta time.Time, status bool) {
	this.Id = id
	this.Nombre = nombre
	this.FechaAlta = fechaAlta
	this.Status = status

}
