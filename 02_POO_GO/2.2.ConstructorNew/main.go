package main

import (
	"fmt"

	"github.com/Lahkpom/myModules/courses"
)

func main() {

	//* En el package de courses ponemos el nombre de la estructura en minúscula para que no
	//* ne tenga un acceso directo. Y se creará un contructor New() que además se podrá
	//* configurar con valores por defecto.

	/*
		c1 := courses.Course{
			Name:  "Prueba",
			Price: 50,
			Classes: []courses.Class{
				{Name: "Clase 1", Description: "Descripción de la clase 1"},
				{Name: "Clase 2", Description: "Descripción de la clase 2"},
				{Name: "Clase 3", Description: "Descripción de la clase 3"},
			},
		}
	*/

	c1 := courses.NewCourse("Prueba", 50, false)

	c1.Classes = []courses.Class{
		{Name: "Clase 1", Description: "Descripción de la clase 1"},
		{Name: "Clase 2", Description: "Descripción de la clase 2"},
		{Name: "Clase 3", Description: "Descripción de la clase 3"},
	}

	c24 := courses.Class{Name: "Clase 4", Description: "Descripción de la clase 4"}
	c25 := courses.Class{Name: "Clase 5", Description: "Descripción de la clase 5"}

	c1.Classes = append(c1.Classes, c24, c25)

	c1.PrintClasses()
	fmt.Printf("El precio del curso es %v\n", c1.Price)

	c1.ChangePrice(100)
	fmt.Printf("El precio del curso actualizado es: %f\n", c1.Price)

	//c2 := courses.NewCourse("Prueba 2", 50)
	//fmt.Println(c2)
}
