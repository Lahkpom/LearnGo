package main

// Las funciones se declaran por fuera de las estructuras y pueden ser reutilizables para todos
// Podemos hacer uso de las estructuras siempre y cuando estén dentro del mismo package

//! CON GO BUILD PODEMOS GENERAR NUESTRO BINARIO PARA EJECUTICÓN

// Podemos tener otros archivos solo con declaración de función aún teniendo la declaración de la estructuraa
// En otro distinto

/*
func (c Course) PrintClasses() {
	fmt.Println("Las clases de este curso son:")
	for _, class := range c.Classes {
		fmt.Println(class)
	}
}

type Course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	Classes map[uint]string
}
*/

// No hace falta importar el contenido del otro archivo, es suficiente con que esté en el mismo package.

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