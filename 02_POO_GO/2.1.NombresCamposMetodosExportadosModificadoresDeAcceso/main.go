package main

import (
	"fmt"

	"github.com/Lahkpom/myModules/courses"
)

//* En go no hay declaraciones como public, private o protected para 
//* indicar la disponibilidad de una función
//* En go existen los identificadores exportados o no exportados.
//* Es decir, indicar qué funciones pueden utilizar los usuarios de 
//* nuestros paquetes.
//* Esto se hace poniendo la primer letra de la función en mayúscula o minúscula
//* Si es mayúscula son exportados y si es minúscula son no exportados.
//! Esto es para funciones y variables
//! El poner las cosas en minúsculas es como un private
//* Esto es el pie para setters y getters

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

	c1.PrintClasses()

	c1.ChangePrice(100)
	fmt.Printf("El precio del curso es: %f\n", c1.Price)
}
