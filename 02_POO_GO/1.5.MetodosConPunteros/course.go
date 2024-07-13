package main

import "fmt"

type Course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	// Classes map[uint]string
	Classes []Class
}

// PrintClasses() Recibe un objeto Course e imprime el nombre y descripción de sus clases
func (c Course) PrintClasses() {
	fmt.Println("Las clases de este curso son:")
	for _, class := range c.Classes {
		fmt.Println(class.Name)
		fmt.Println(class.Description)
		fmt.Println("==================================")
	}
}

// ChangePrice() Recibe un valor y lo actualiza en el precio
func (c *Course) ChangePrice(newPrice float64) {
	c.Price = newPrice
}

//! CUANDO NECESITAMOS ACTUALIZAR ALGO DE LA ESTRUCTURA USAMOS EL TIPO PUNTERO
//! SI NO NECESITAMOS ACTUALIZAR NADA SOLO USAMOS LA COPIA
