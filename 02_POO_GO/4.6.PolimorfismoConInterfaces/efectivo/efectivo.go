package efectivo

import (
	"fmt"
)

type Efectivo struct{}

func New() *Efectivo {
	return &Efectivo{}
}

func (e *Efectivo) Pagar() {
	fmt.Println("Pagando con Efectivo")
}
