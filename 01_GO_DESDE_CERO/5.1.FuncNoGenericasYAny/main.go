package main

import "fmt"

func main() {
	printList("uno", "dos", "tres")
	fmt.Println("")

	//* Si quisiera imprimir una lista de int, no podría usar la misma que printList
	//* Tendría que duplicar la misma lógica para cada tipo de dato

	//* La llegada de las funciones genéricas trae la posibilidad de dar a la función parámetreos de cualquier tipo
	//* con el comando any
	printListAny("uno", "dos", "tres")
	fmt.Println("")
	printListAny(1, 2, 3)
	fmt.Println("")
	//* Permite incluso que se mezclen los datos
	printListAny(1, "dos", 3)

	aux := []string{"uno", "dos", "tres"}
	aux2 := []int{1, 2, 3}
	//* Haciendo pruebas vemos que también sirve para otros tipos de dato
	aux3 := []any{1, "dos", 3}

	aux4 := map[any]any{
		1:     "uno",
		"dos": 2,
		3:     "tres",
	}

	fmt.Println(aux)
	fmt.Println(aux2)
	fmt.Println(aux3)
	fmt.Println(aux4)

	fmt.Println("A partir de acá")
	printListAny(aux, aux2, aux3, aux4)
}

// * Esta es una función no genérica, una función normal
func printList(list ...string) {
	for _, v := range list {
		fmt.Println(v)
	}
}

// * Any es un nuevo tipo de dato
func printListAny(list ...any) {
	for _, v := range list {
		fmt.Println(v)
	}
}
