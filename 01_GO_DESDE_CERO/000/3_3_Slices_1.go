package main

import "fmt"

func main() {
	//*Son punteros a un array, no guardan datos
	//*Se le pasa un rango del Array y almacena solo las direcciones de lo indicado
	things := [7]string{"peugeot", "chevrolet", "vw", "patrulla", "bomberos", "ferrari", "globo"}
	cars := things[0:6] //*Siempre hay que poner uno demas porque el ultimo no es incluyente
	red := things[4:7]

	//*Para achicar la escritura, podemos evitar poner el primer numero y el ultimo
	cars2 := things[:6]
	red2 := things[4:]

	fmt.Println(things)
	fmt.Println(cars)
	fmt.Println(red)
	fmt.Printf("%T, %s\n", cars, cars)
	red[0] = "palanca" //Al estar trabajando con punteros, si se cambia en uno se cambia en todos
	fmt.Println(things)
	fmt.Println(cars2, red2)
}
