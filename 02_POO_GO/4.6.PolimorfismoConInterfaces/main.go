package main

//* Acá mostramos cómo usar el polimorfismo con interfaces

import (
	"main/fabrica"
)

func main() {
	metodo := fabrica.GetMetodo("efectivo")
	metodo.Pagar()
}
