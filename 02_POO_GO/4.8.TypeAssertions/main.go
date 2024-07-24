package main

import (
	"fmt"
	"strings"
)

func Wrapper(i interface{}) {
	fmt.Printf("Valor: %v, Tipo: %T\n", i, i)

	// Cómo podemos controlar qué tipo de dato específico estamo tratando
	// para poder darle un tratamiento según lo deseado

	// Se pone i.(type) donde type es el tipo que deseo controlar
	// ok indica si es true lo que yo teoy queriendo poner
	// v almacena el valor de la interface ingresada por parámetro SOLO si ok == true
	
	v, ok := i.(int) 

	if ok {
		v = 9999
		fmt.Println(v)
	}
	
	y, ok1 := i.(float64) 
	
	if ok1 {
		y = 999.999
		fmt.Println(y)
	}
	

	z, ok2 := i.(string) 
	if ok2 {
		strings.ToUpper(z)
		fmt.Println(z)
	}
	

	fmt.Println("####################")
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


