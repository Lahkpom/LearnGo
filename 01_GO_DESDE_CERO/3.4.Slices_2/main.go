package main

import "fmt"

func main() {
	//* len() # de elementos del slice
	//* cap() # de elementos del array original a partir del índice desde donde se creó
	aAnimals := [...]string{"Monkey", "Dog", "Cat", "Bird", "Elephant"}
	sPets := aAnimals[1:3]

	fmt.Println("Animals:", aAnimals)
	fmt.Println("Pets:", sPets)
	fmt.Println("len pets:", len(sPets))
	fmt.Println("cap pets:", cap(sPets))
	fmt.Println("len animals:", len(aAnimals))
	fmt.Println("cap animals:", cap(aAnimals))

	//* A diferencia del array, en el slice se puede agregar elementos con la funcion slice = append(slice, valor.1, valor.n)
	//* Se pueden agregar elementos sin límite, misntras no sobrepace la capacidad del array va a modificar el array original.
	//* Al momento en que sobrepasemos la capacidad del array original, automáticamente se creará un nuevo array a donde
	//* ahora va a apuntar el slice, y el anterior lo dejará intacto. Este nuevo array es uno diferente que duplica la capacidad
	//* del array anterior, manteniendo la información que hasta ahora tenía
	sPets = append(sPets, "Parrot", "Parrot1", "Parrot2", "parrot3")

	fmt.Println("Animals:", aAnimals)
	fmt.Println("Pets:", sPets)
	fmt.Println("New len pets:", len(sPets))
	fmt.Println("New cap pets:", cap(sPets))

	//* También podemos crear slices sin hacer referencia a un array existente
	//* Con esta syntaxis, sin especificar la cantidad en []
	sPets2 := []string{}

	fmt.Println("Pets:", sPets2)
	fmt.Println("New len pets:", len(sPets2))
	fmt.Println("New cap pets:", cap(sPets2))

	sPets2 = append(sPets2, "pet1")
	fmt.Println("New len pets:", len(sPets2))
	fmt.Println("New cap pets:", cap(sPets2))

	sPets2 = append(sPets2, "pet2", "pet3", "pet4")
	fmt.Println("New len pets:", len(sPets2))
	fmt.Println("New cap pets:", cap(sPets2))

	//* De esta forma no trabaja con espacios y capacidades de un array existente, sino que creamos un nuevo array que si permite
	//* agregar más elementos y actualiza su tamaño y capacidad constantemente

	//* Otra forma de crear slices es con la función make(type, len, cap) donde yo al momento de iniciar el slices le indico
	//* el tamaño y capacidad que quiero que tenga desde un inicio. Si el len supera el cap, automática se duplica el cap

	sPets3 := make([]string, 0, 3)
	fmt.Println("New make len pets:", len(sPets3))
	fmt.Println("New make cap pets:", cap(sPets3))

	sPets3 = append(sPets3, "pet2", "pet3", "pet4")
	fmt.Println("New make len pets:", len(sPets3))
	fmt.Println("New make cap pets:", cap(sPets3))

	sPets3 = append(sPets3, "pet5")
	fmt.Println("New make len pets:", len(sPets3))
	fmt.Println("New make cap pets:", cap(sPets3))

	//! CON CADA DUPLICACIÓN QUE GO HACE POR EL SLICE SE TOMA EL TRABAJO DE CREAR UN NUEVO ARRAY

	//* Los slice tienen como valor cero nil que es null (imagino que lo array también)
	var sSlice []string
	fmt.Printf("Type = %T, Value = %v, Value2 = %s, Value3 = %q, Dir = %v\n", sSlice, sSlice, sSlice, sSlice, &sSlice)
	fmt.Println(sSlice == nil) //* Esto va a dar true

	//* En cambio si usamos el operador de asignación de variable corta, estaríamos inicializando el slice, por lo que no sería nil sino empty
	sSlice2 := []string{}
	fmt.Printf("Type = %T, Value = %v, Value2 = %s, Value3 = %q, Dir = %v\n", sSlice2, sSlice2, sSlice2, sSlice2, &sSlice2)
	fmt.Println(sSlice2 == nil) //* Esto va a dar true

}
