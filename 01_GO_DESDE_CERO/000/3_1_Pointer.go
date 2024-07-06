package main

import "fmt"

func main() {
	//*El puntero nos indica la dirección de memoria donde se almacena nuestro dato de la variable
	//*se hace con el Printf, y anteponiendo el & a la varable
	var a string = "Hola"

	fmt.Printf("Dirección: %v\n", &a)

	//*Para alamacenar un puntero en una variable, se debe anteponer un * al tipo de dato
	//*De esta forma indicamos que lo almacenado no es un valor sino una dirección
	var pointer *string = &a
	fmt.Println(pointer)
	fmt.Printf("Dirección: %v\n", pointer)
	fmt.Println(*pointer) //*De esta forma accedemos al valor del puntero, la inversa

	//*Si quisieramos modificar el valor de la variable desde el puntero
	*pointer = "Chau"
	fmt.Println(pointer)
	fmt.Println(*pointer)
}
