package main

import "fmt"

func main() {
	//Las variables se declaran con la palabra reservada var
	//seguido del nombre de la variable, y a continuación el tipo
	var apple string
	//en este caso un string, siempre va entre comillas dobles
	apple = "manzana"
	fmt.Println(apple)

	//podemos agrupar las variables en un solo var
	/*
		var banana string = "banana"
		var orange string = "naranja"
		var pineapple string = "ananá"
	*/

	var (
		banana    string = "banana"
		orange    string = "naranja"
		pineapple string = "ananá"
	)

	//Podemos agrupar por tipo de dato en una misma línea
	/*
		var banana, orange, pineapple string = "banana", "naranja", "ananá"
	*/

	fmt.Println(apple, banana, orange, pineapple)

	//*También podemos saltearnos la palabra var
	//*y el tipo de dato, aplicando el := operador de asignación de variable corta
	//*que infiere el tipo de dato de acuerdo al valor asignado
	//*una vez inicializada la variable ya no puede cambiar su tipo de dato

	watermelon := "sandía"
	fmt.Println(watermelon)

	/*
		No podemos volver a utilizar el := una vez quew ya fue inicializada la variable
	*/
	//watermelon = 1
	watermelon = "asd"
	fmt.Println(watermelon)

	//Podemos agrupar independientemente del tipo de dato en una misma línea aplicando el asigandor corto
	/*
		banana, orange, pineapple := "banana", 1, true
	*/

	// En este caso estamos combinando la reasignación de valor con una nueva declaración en la misma línea
	watermelon, lemon := "dsa", "limón"
	fmt.Println(watermelon, lemon)

	pear, lemon := "Pera", "limoncito"
	fmt.Println(pear, lemon)
}
