package main

import (
	"fmt"
)

// Todos implementan una interface vacía ya que no requiere cumplir ninguna función
// como requisito, por lo tanto cualquier cosa que se cree la contiene
// Esto nos permite manipular con funciones ingresos que no sabes con seguridad
// de qué tipo de dato van a ser
func Wrapper(i interface{}) {
	fmt.Printf("Valor: %v, Tipo: %T\n", i, i)
}

func main() {
	Wrapper(12)
	Wrapper(12.5)
	Wrapper(-50)
	Wrapper(true)
	Wrapper("hola")
	Wrapper(Person{})
}

type Person struct {}


