package main

import (
	"fmt"
	"strings"
)

func main() {
	sayHello()
	print("Hello world 2")
	txt := "Leonel"
	fmt.Println(txt)
	toUpper(txt)
	fmt.Println(txt)

	//* Al ser un puntero hay que mandarlo con el &
	toUpper2(&txt)
	fmt.Println(txt)

}

//* Estructuras básicas de las funciones (sin retorno)
//* func nameFunc(var1 varType, varN varType) {
//*		body
//* }
//* Si tenemos varios parámetros, podemos agruparlas por tipo
//* (a, b, c string, d, e, f int)

func sayHello() {
	fmt.Println("Hello world")
}

func print(a string) {
	fmt.Println(a)
}

func toUpper(txt string) {
	txt = strings.ToUpper(txt)
	fmt.Println(txt)
}

//* Estos son parámetros por valor, lo que quiere decir que todo lo que pasa dentro de la función se hace en copias del
//* original (NO SE MODIFICA EL ORIGINAL)
//* Para poder modificar el original neceitamos usar valores por referencia o puntero
//* Se le pone el * al tipo de dato por parámetro

func toUpper2(txt *string) {
	//* Usando el puntero es que modificamos la variable original
	*txt = strings.ToUpper(*txt)
}
