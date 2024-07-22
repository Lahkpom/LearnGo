package main

import (
	"fmt"
)

type Storager interface {
	Get() string
	Set(string)
}

type Person struct {
	name string
}

func NewPerson(name string) *Person {
	return &Person{name}
}

func (p *Person) Get() string {return p.name}

// Recordar que para actualizar un valor debo tener un tipo puntero
// Sino solo se actualizará la copia
func (p *Person) Set(name string) {p.name = name}

// Storager sería cualquier variable o estructura que cumpla con todos
// los métodos que solicita Storager (En este caso Person lo cumple)
func Exec(s Storager, name string) {
	s.Set(name)
	fmt.Println(s.Get())
}

func main() {
	p1 := NewPerson("Leo")
	fmt.Println(p1.Get())
	p1.Set("Leonel")
	fmt.Println(p1.Get())
	fmt.Println("####")
	// Le pasamos un puntero 
	Exec(p1, "Otro gato")
}
