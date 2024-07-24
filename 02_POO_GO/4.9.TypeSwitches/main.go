package main

import (
	"fmt"
	"strings"
)

func Wrapper(i interface{}) {
	fmt.Printf("Valor: %v, Tipo: %T\n", i, i)
	
	// Para estos casos donde necesitamos validar más de un assertion podemos simplificarlo
	// usando el type switches

	switch v := i.(type) {
	case int:
		fmt.Printf("%v es de tipo %T\n", v, v)
	case float64:
		fmt.Printf("%v es de tipo %T\n", v, v)
	case string:
		fmt.Printf("%v es de tipo %T, mayus = %v\n", v, v, strings.ToUpper(v))
	case Person:
		fmt.Printf("%v es de tipo %T\n", v, v)
	}
	fmt.Println("#####")
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


