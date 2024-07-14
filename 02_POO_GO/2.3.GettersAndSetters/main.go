package main

import (
	"fmt"

	"github.com/Lahkpom/myModules/courses"
)

func main() {

	//* Lo que es Getters y Setters lo apliqué en el paquete courses poniendo las variables 
	//* con minúscla

	c1 := courses.NewCourse("Prueba", 50, false)

	c1.AddClass("Clase 1", "Descripción de la clase 1")
	c1.AddClass("Clase 2", "Descripción de la clase 2")
	c1.AddClass("Clase 3", "Descripción de la clase 3")

	c1.AddClass("Clase 4", "Descripción de la clase 4")
	c1.AddClass("Clase 5", "Descripción de la clase 5")

	c1.PrintClasses()
	fmt.Printf("El precio del curso es %v\n", c1.GetPrice())

	c1.ChangePrice(100)
	fmt.Printf("El precio del curso actualizado es: %f\n", c1.GetPrice())

}
