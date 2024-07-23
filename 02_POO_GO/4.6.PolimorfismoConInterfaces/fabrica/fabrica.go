package fabrica

import (
	"main/efectivo"
	"main/metodoPago"
	"main/paypal"
)

type Fabrica struct{}

// GetMetodo() Devuelve un MetodoPago de tipo interface, 
// osea que puede tomar la forma de cualquier estructura que tenga implementada esa interfaz
func GetMetodo(metodo string) metodoPago.MetodoPago {
	switch metodo {
	case "efectivo":
		return efectivo.New()
	case "paypal":
		return paypal.New()
	default:
		return nil
	}
}
