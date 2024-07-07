package main

import (
	"fmt"
)

type Produc struct {
	id          uint
	description string
	price       float64
}

//* Con la formaque se ve arriba, tendríamos un problema cuando necesitemos que la información se comparta en
//* dos bases de dato distintas y le cambio algún tipo de dato, tendríamos que duplicar el código.

type ProducV2 struct {
	id          string
	description string
	price       float64
}

//* Para solucionar esto tenemos los genéricos que nos permiten especificar el tipo de dato que debe recibir

type ProductFinal[T uint | string] struct {
	id          T
	description string
	price       float64
}

func main() {
	// p1 := Produc{1, "Notebook", 1800}
	// p1V2 := Produc{1, "Notebook", 1800}
	// p2 := Produc{2, "Impressora", 800}
	// p2V2 := Produc{2, "Impressora", 800}
	// p3 := Produc{3, "Mouse", 80}
	// p3V2 := Produc{3, "Mouse", 80}

	p1 := ProductFinal[uint]{1, "Notebook", 1800}
	p2 := ProductFinal[string]{"2", "Impressora", 800}
	p3 := ProductFinal[uint]{3, "Mouse", 80}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)

	products := []ProductFinal[uint]{p1, p3}

	for _, p := range products {
		fmt.Println(p.id, p.description, p.price)
	}
}
