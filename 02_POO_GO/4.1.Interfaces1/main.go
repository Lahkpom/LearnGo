package main

import (
	"fmt"
)

// Las interfaces se escriben con el sufijo er 
// Se le deben poner la firma de métodos que luego se implementarán
// en aquellas estructuras que la utilicen
type Greeter interface {
	Greet()
}

// No es necesario escribirlo explícitamente en la definición de la struct
// Solo con crear una func llamada Greet ya Go entiende que esa struct
// impletmenta esa interface

type Person struct {
	name string
}

func NewPerson(name string) Person {
	return Person{name}
}

func (p Person) Greet() {
	fmt.Printf("Hi! I'm, %v\n", p.name)
}

type Text string

func (t Text) Greet() {
	fmt.Printf("Hi! Text = %v\n", t)
}

func main() {
	p1 := NewPerson("Leonel")
	p1.Greet()

	var t1 Text = "Vas a morir Moe, WIIIII"
	t1.Greet()

	// También podemos crear una variable directa de la interface
	var g1 Greeter = Person{"Facundo"}
	var g2 Greeter = Text("Un gran poder conlleva una gran responsabilidad")
	g1.Greet()
	g2.Greet()
}
