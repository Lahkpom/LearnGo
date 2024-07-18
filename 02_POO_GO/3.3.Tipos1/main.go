package main

import "fmt"

// El conjunto de métodos que utilizan estos nuevos type
// deben ser declarados en el mismo archivo en qiue se 
// declara el type
// * Es por esto que no podemos declara métodos para los type
// * predeterminados como string bool o los numéricos ya que no
// * tenemos acceso a los archivos donde fueron creados
type numero int

// Podemos crear alias que hagan referencia a otros tipos
// * No se para que sirve pero se puede
type course struct {
	name string
}

type myAlias = course

// Podemos crear una definición de type que si bien hace referencia
// al type original y comparte sus campos, NO comparte sus métodos
// y requiere métodos propios, arranca con un conjunto de métodos vacío

type newCourse course

func main() {
	// Puedo declarar nuevos type en base a los existentes
	var n numero = 5
	fmt.Println(n)
	n.Print()
	fmt.Printf("Type: %T\n", n)

	c := myAlias{name: "Leo"}
	c.Print()
	fmt.Printf("myAlias type: %T\n", c)

	nc := newCourse{name: "Facu"}
	fmt.Printf("nc type: %T\n", nc)

}

func (n numero) Print() {
	fmt.Println(n)
}

func (c course) Print() {
	fmt.Println(c)
}
