package main

import "fmt"

func main() {
	//*Para declarar la variable ponemos:
	//*var nombreVar tipoVar

	var a string
	a = "Hola 😳"
	fmt.Println(a)

	var b string = "Hola dos"
	fmt.Println(b)

	//*Si tenemos muchas variables del mismo tipo, se pueden agrupar
	var (
		c string = "Hola tres"
		d string = "Hola cuatro"
		e string
	)
	e = "Hola cinco"
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	//*Asignación de Variable Corta
	//*Se le asigna a la variable el tipo segun el dato que se le ingrese
	//*Igualmente, una vez asignado el tipo de dato a la variable, solo mantendra ese mismo tipo
	apple := "Manzana"
	verdad := true
	fmt.Println(apple, verdad)

}
