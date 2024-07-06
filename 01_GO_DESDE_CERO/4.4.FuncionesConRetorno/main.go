package main

import (
	"fmt"
	"strings"
)

func main() {
	aux := sum(2, 8)
	fmt.Println(aux)
	fullName := "Leonel Alejandro Hidalgo"

	//* Creamos dos variables para recibir cada valor que devuelve la función
	a, b := conver(fullName)
	fmt.Println(a)
	fmt.Println(b)

	a, b = conver2(fullName)
	fmt.Println(a)
	fmt.Println(b)
}

// * Hay que aclarar el valor de retorno
func sum(a, b int) int {
	return a + b
}

// * Las funciones pueden retornar dos valores, se deben poner los tipos entre paréntesis
func conver(txt string) (string, string) {
	a := strings.ToLower(txt)
	b := strings.ToUpper(txt)
	return a, b
}

// * Puedo declarar lsa variables a retornar dentro del paréntesis
// * En este caso no debo utilizar el operador de asignación de variable corta
// ! Se recomienda hacerlo cuando la función es pequeña, no usar para funciones complejas
func conver2(txt string) (a string, b string) {
	a = strings.ToLower(txt)
	b = strings.ToUpper(txt)
	//* Tampoco es necesario especificar nada en el return, devolverá las variables que declaramos en el
	//* paréntesis y en el mismo orden
	return
}
