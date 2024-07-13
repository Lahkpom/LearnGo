package main

import (
	"fmt"

	"github.com/Lahkpom/myModules/courses"
)

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
	c1 := courses.Course{
		Name:  "Prueba",
		Price: 50,
		Classes: []courses.Class{
			{Name: "Clase 1", Description: "Descripción de la clase 1"},
			{Name: "Clase 2", Description: "Descripción de la clase 2"},
			{Name: "Clase 3", Description: "Descripción de la clase 3"},
		},
	}

	c24 := courses.Class{Name: "Clase 4", Description: "Descripción de la clase 4"}
	c25 := courses.Class{Name: "Clase 5", Description: "Descripción de la clase 5"}

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
