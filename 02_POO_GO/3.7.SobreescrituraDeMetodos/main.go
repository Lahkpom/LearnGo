package main

import (
	"fmt"
)

type Person struct {
	name string
	age uint8
}

func NewPerson(name string, age uint8) Person {
	return Person{name, age}
}

func (p Person) Greet() {
	fmt.Printf("Hello, I'm %v, and I'm %v\n", p.name, p.age)
}

// Para embeber un type en una estructura la debemos poner dentro
// de esta forma podríamos decir que hereda o es subclase de la anterior
// ya que se pueden utilizar sus métodos y atributos
// no es person Person, es solo Person
type Employee struct{
	Person
	salary float64
}

// Primero hacemos ingresar lo del Person y después lo propio  del Employee
func NewEmployee(name string, age uint8, salary float64) Employee {
	return Employee {
		NewPerson(name, age),
		salary,
	}
}

func (e Employee) Payroll() {
	fmt.Printf("Payroll: %v\n", e.salary * 0.90)
}

// Esto no es técnicamente sobre escritura
// En Go se le dice ocultación de métodos
// Lo que hace Go es revisar si Employee tiene una función Greet()
// En caso que no la tenga, revisa si alguna de las estructuras 
// embebidas tiene una
// No es sobre escritura porque todavía podría acceder a la función 
// Greet() de la embebida
// Al momento de llamar a la variable con la que creemos Employee
// le ponemos .Embebida, y ahí Greet()

func (e Employee) Greet() {
	fmt.Println("Hello from Employee")
}

func main()  {
	p1 := NewPerson("Leo", 25)
	p1.Greet()

	e1 := NewEmployee("Fiore", 23, 100)
	e1.Greet()
	e1.Payroll()
	// Así se llama a la función oculta de la embebida
	e1.Person.Greet()
}
