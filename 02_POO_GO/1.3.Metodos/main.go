package main

import "fmt"

type Course struck {
	Name string
	Price float64
	IsFree bool
	UserIDs []uint
	Classes map[uint]string
}

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
	c.PrintClasses()
}
