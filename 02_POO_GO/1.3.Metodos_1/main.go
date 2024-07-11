package main

import "fmt"

type Course2 struct {
	Name string
	Price float64
}

type Course struct {
	Name string
	Price float64
	IsFree bool
	UserIDs []uint
	Classes map[uint]string
}

// Al poner en una func el paréntesis adelante indicando que se recibe una
// variable de tipo <struct>, y luego el nombre de la func, se le asignará
// esa función a la struct y se podrá llamar a la función poniendole un . a 
// la variables creada

func (c Course) PrintClasses() {
	fmt.Println("Las clases de este curso son:")
	for _, class := range c.Classes {
		fmt.Println(class)
	}
}

func main() {
	c := Course{
		Name: "Prueba", 
		Classes: map[uint]string{
			1: "Clase 1", 
			2: "Clase 2",
		},
	}
	// En vez de PrintClasses(c)
	c.PrintClasses()
}
