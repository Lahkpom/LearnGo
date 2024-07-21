package main

import (
	"fmt"
)

type Greeter interface {
	Greet()
}

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

// Los valores de las interfaces se consideran una 
// tupla entre un valor y un tipo concreto 

func GreetAll(gs ...Greeter){
	for _, g := range gs {
		g.Greet()
		fmt.Printf("\t- Value: %v. Type: %T\n", g, g)
	} 
}

func main() {
	p1 := NewPerson("Leonel")
	p1.Greet()

	var t1 Text = "Vas a morir Moe, WIIIII"
	t1.Greet()

	var g1 Greeter = Person{"Facundo"}
	var g2 Greeter = Text("Un gran poder conlleva una gran responsabilidad")
	g1.Greet()
	g2.Greet()
	
	// Acá podemos ver que podemos mandar incluso las instancias directas
	// de PErson y Text ya que ambas implementan Greet() y son Greeters
	fmt.Println("################")
	GreetAll(p1, t1, g1, g2)
}
