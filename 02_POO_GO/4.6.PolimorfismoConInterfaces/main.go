package main

//* Acá mostramos cómo usar el polimorfismo con interfaces

import (
	"fmt"
	"main/fabrica"
	"main/metodoPago"
)

const (
	PP = "PAYPAL"
	EF = "EFECTIVO"
)

func main() {
	var metodo string
	fmt.Println("Métodos de pago:")
	fmt.Printf("\t1. %v\n", PP)
	fmt.Printf("\t2. %v\n", EF)
	fmt.Printf("Ingrese la opción deseada: ")
	_, err := fmt.Scanln(&metodo)
	
	for {
		if err == nil {
			break
		}
		fmt.Println("Inválido! Ingréselo nuevamente: ")
		_, err = fmt.Scanln(&metodo)
	}

	/*
	for err != nil {
		fmt.Println("Inválido! Ingréselo nuevamente: ")
		_, err = fmt.Scanln(&metodo)
	}
	*/

	if metodo == "" {
		panic("Está vacío")
	}

	var m metodoPago.MetodoPago

	for{
		m = fabrica.GetMetodo(metodo)
		if m != nil {
			break
		}
		fmt.Printf("El método ingresado es incorrecto, ingréselo nuevamente: ")
		_, err = fmt.Scanln(&metodo)
		if err != nil {
			panic("LA CONCHA DE TU MADRE!")
		}
	}
	m.Pagar()
}
