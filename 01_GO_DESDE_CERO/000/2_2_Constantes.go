package main

import "fmt"

func main() {
	//*Se debe asignar el valor al mismo momento en que se declara
	//*Se reemplaza la àlabra var por const
	const a string = "a"
	fmt.Println(a)

	//*Cuando utilizamos variables/constantes que son incrementales
	//*se las puede agrupar y usar una funcion iota que incrementa
	//*automaticamente el valor de cada variable
	//*iota siempre vale 0
	const (
		Jan = iota + 1
		Feb
		Mar
		Apr
		May
		Jun
		Jul
	)

	fmt.Println(May)
}
