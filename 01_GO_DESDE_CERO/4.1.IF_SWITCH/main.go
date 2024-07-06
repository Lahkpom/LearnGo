package main

import "fmt"

func main() {
	//! IF
	//* La condición no va entre paréntesis
	if true {
		fmt.Println("t")
	} else if false {
		fmt.Println("f")
	} else {
		fmt.Println("x")
	}

	//* En go el if permite crear una variable dentro del mismo ciclo que sea solo de uso interno

	if aux := false; aux {
		fmt.Println("t")
	} else if !aux {
		fmt.Println("f")
	} else {
		fmt.Println("x")
	}

	//! SWITCH
	//* Tampoco es necesario poner las cosas en paréntesis
	aux := 2
	switch aux {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("Default")
	}

	//* Go permite que en los casos en que la respuesta sea la misma, se agrupen
	//* Se separan con coma
	aux2 := 2
	switch aux2 {
	case 1, 2:
		fmt.Println("1")
	case 3, 4:
		fmt.Println("2")
	default:
		fmt.Println("Default")
	}

	//* Otra alternativa es no aclarar la expresión después del switch y ponerla directamente en el caso
	switch {
	case aux2 == 1 || aux2 == 2:
		fmt.Println("1")
	case aux2 == 3 || aux2 == 4:
		fmt.Println("2")
	default:
		fmt.Println("Default")
	}

}
