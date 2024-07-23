package paypal

import (
	"fmt"
)

type PayPal struct{}

func New() *PayPal {
	return &PayPal{}
}

func (p *PayPal) Pagar() {
	fmt.Println("Pagando con PayPal")
}
