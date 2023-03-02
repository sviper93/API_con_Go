package model // Es una buena práctica crear dentro de un paquete un archivo que se llame igual
// y en este archivo crear las constantes, interfaces, las variables que son globales
import "errors"

var ( // creo las variables que almacenarán los errores más comunes
	// ErrPersonCanNotBeNil la persona no puede ser nula
	ErrPersonCanNotBeNil = errors.New("la persona no puede ser nula")
	// ErrIDPersonDoesNotExists la persona no existe
	ErrIDPersonDoesNotExists = errors.New("la persona no existe")
)
