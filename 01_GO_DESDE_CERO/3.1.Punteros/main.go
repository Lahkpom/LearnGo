package main

import "fmt"

func main() {
	//* anteponiendo el & antes del nombre de la variable, nos da la dirección en memoria donde se almacena
	//* para indicarle a go que lo que vamos a almacenar es una dirección, ponemos * antes del tipo de dato

	var color string = "Rojo"

	fmt.Printf("Type = %T, Value = %v, Value2 = %s, Dir = %v\n", color, color, color, &color)

	//* dir := &color
	var dir *string = &color

	//* La desferenciación es obtener el valor de la variable a la que apunta
	//* El value es el calor que guarda con la Dir a la otra variable
	//* El Dir es la dirección de la propia variable, no de la que apunta
	fmt.Printf("Type = %T, Value = %v, Dir = %v, Desfereciación = %v\n", dir, dir, &dir, *dir)

	dir2 := dir

	dir3 := dir2

	fmt.Printf("Type = %T, Value = %v\n", dir3, dir3)

	//* También podemos modificar el valor de la variable original desde el puntero
	*dir3 = "Azul"
	fmt.Printf("Type = %T, Value = %v, Dir = %v\n", color, color, &color)
	fmt.Printf("Type = %T, Value = %v, Dir = %v, Desfereciación = %v\n", dir3, dir3, &dir3, *dir3)

}
