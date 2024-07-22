package main

import (
	"fmt"
)

type Greeter interface {
	Greet()
}

type Byer interface {
	Bye()
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

func (p Person) Bye() {
	fmt.Println("See you later!")
}

type Text string

func (t Text) Greet() {
	fmt.Printf("Hi! Text = %v\n", t)
}

func (t Text) Bye() {
	fmt.Println("Bye func.")
}

// Los valores de las interfaces se consideran una 
// tupla entre un valor y un tipo concreto 

// Esta función queda inutil pero la dejamospara los ejemplos
func GreetAll(gs ...Greeter){
	for _, g := range gs {
		g.Greet()
		fmt.Printf("\t- Value: %v. Type: %T\n", g, g)
	} 
}


// Es posible embeber en insterfaces pero solo pueden ser otras
// interfaces, además es requisito que no haya colición 
// entre las interfaces por tener un nombre idéntico

type GreeterByer interface {
	Greeter
	Byer
} 

func All(gbs ...GreeterByer) {
	for _, gb := range gbs {
		// Podemos llamar a ambas
		gb.Greet()
		gb.Bye()
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
	fmt.Println("################")
	p1.Bye()
	t1.Bye()
	// Estos no pueden implementarla porque son directos del insterface 
	// Greeter
	// g1.Bye()
	// g2.Bye()
	fmt.Println("################")
	All(p1, t1)
}
