package main

import "fmt"

func main() {
	var days [7]string

	days[0] = "Monday"
	days[1] = "Tuesday"
	days[2] = "Wednesday"
	days[3] = "Thursday"
	days[4] = "Friday"
	days[5] = "Saturday"
	days[6] = "Sunday"

	fmt.Println(days)
	fmt.Printf("Tipo: %T, Valor: %v\n", days, days)

	//*Forma corta de inicializar
	months := [6]string{"January", "February", "March", "April", "May", "June"}
	fmt.Println(months)

	//*Para poder generar un array sin especificar el tamaño, se pone ... en los corchetes
	//*De este modo el array toma el tamaño según la cantidad de valores que posee
	things := [...]uint8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("Tipo: %T, Valor: %v\n", things, things)

}
