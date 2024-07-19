package main

import (
	"fmt"
)

// Podemos crear types basados en los type base

type newBool bool

// A estos les podemos crear nuevas funciones que necesitemos y
// no vengan por defecto

func main() {
	var aux newBool = true
	fmt.Println(aux.ToString())
	aux = false
	fmt.Println(aux.ToString())
}

func (nb newBool) ToString() string {
	if nb {
		return "TRUE"
	}
	return "FALSE"
}
