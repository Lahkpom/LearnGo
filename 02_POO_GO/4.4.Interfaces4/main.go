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

type GreeterByer interface {
	Greeter
	Byer
} 

func All(gbs ...GreeterByer) {
	for _, gb := range gbs {
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
	
	fmt.Println("################")
	p1.Bye()
	t1.Bye()
	
	fmt.Println("################")
	All(p1, t1)
}
