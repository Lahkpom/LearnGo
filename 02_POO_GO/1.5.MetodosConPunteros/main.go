package main

import "fmt"

// Las funciones se declaran por fuera de las estructuras y pueden ser reutilizables para todos
// Podemos hacer uso de las estructuras siempre y cuando estén dentro del mismo package

//! CON GO BUILD PODEMOS GENERAR NUESTRO BINARIO PARA EJECUTICÓN

// Podemos tener otros archivos solo con declaración de función aún teniendo la declaración de la estructuraa
// En otro distinto

/*
func (c Course) PrintClasses() {
	fmt.Println("Las clases de este curso son:")
	for _, class := range c.Classes {
		fmt.Println(class)
	}
}

type Course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	Classes map[uint]string
}
*/

// No hace falta importar el contenido del otro archivo, es suficiente con que esté en el mismo package.

func main() {
	c1 := Course{
		Name:  "Prueba",
		Price: 50,
		Classes: []Class{
			{"Clase 1", "Descripción de la clase 1"},
			{"Clase 2", "Descripción de la clase 2"},
			{"Clase 3", "Descripción de la clase 3"},
		},
	}

	c24 := Class{"Clase 4", "Descripción de la clase 4"}
	c25 := Class{"Clase 5", "Descripción de la clase 5"}

	c1.Classes = append(c1.Classes, c24, c25)

	// En vez de PrintClasses(c)
	c1.PrintClasses()

	//* Para poder pasar un puntero debo poner el objeto entre () y con el & por delante
	//* Sino solo vamos a trabajr con la copia y el valor real nunca se actualiza
	// (&c1).ChangePrice(100)
	// Al final dijo que si bien esta sería la sintaxis completa, go pasa el puntero o desferenciación
	// automáticamente
	c1.ChangePrice(100)
	fmt.Printf("El precio del curso es: %f\n", c1.Price)
}
