package main

import "fmt"

func main() {
	//* Para declarar un Array
	//* var nombreArray [size]type
	//* Los array son de tamaño fijo, una vez declarado no se puede modificar
	var aCountries [3]string

	//* Para asignarle valores hay que indicar el índice
	aCountries[0] = "Argentina"
	aCountries[1] = "Perú"
	aCountries[2] = "Francia"

	fmt.Println(aCountries)

	//! Arrays de asignación directa
	aNames := [4]string{"Leo", "", "Facu"}

	aSurnames := [4]string{
		"Hidalgo",
		"Jara",
		"Flores",
		"Naranjo",
	}

	aSurnames2 := [4]string{
		"Hidalgo",
		"Jara",
		"Flores",
		"Naranjo"} //* Así nos ahorramos la última coma

	fmt.Printf("%s, %q\n", aNames[0], aNames[0])
	fmt.Printf("%s, %q\n", aNames[1], aNames[1])
	fmt.Printf("%s, %q\n", aNames[2], aNames[2])
	fmt.Printf("%s, %q\n", aNames[3], aNames[3])
	fmt.Println(aSurnames)
	fmt.Println(aSurnames2)

	//* Al momento de definir el array con asignación directa, para no poner el número límite nosotros
	//* podemos usar [...], de este modo el tamaño se autodefinirá acorde a lo que le asignemos en {}
	//* Ese tamaño no puede modificarse
	aNames2 := [...]string{"Leo", "Facu", "Fiore"}
	fmt.Println(aNames2)
}
