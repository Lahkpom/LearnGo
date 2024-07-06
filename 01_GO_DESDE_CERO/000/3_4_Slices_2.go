package main

import "fmt"

func main() {
	aAnimals := [5]string{"mono", "perro", "gato", "pájaro", "elefante"}
	sPets := aAnimals[1:3]

	//sPets = append(sPets, "cocodrilo", "iguana", "llama", "vaca")

	fmt.Println(aAnimals)
	fmt.Println("Tamaño total del Array: ", len(aAnimals))
	fmt.Println(sPets)
	fmt.Println("Cantidad de elementos: ", len(sPets))
	fmt.Println("Tamaño total del slice: ", cap(sPets))
	//*La capacidad de un slice se cuenta desde la posición original de su primer elemnto
	//Si el original tiene 5 elementos, y yo hago un slice desde el 1, entonces
	//la capacida de ese slice es de 4

	//*Para aregar elementos desde el slice debemos usar una funcion llamada append()
	//*pasandole por parametro el slice y el elemento unevo
	//*Pero esto modificará tambien el Array original

	sPets = append(sPets, "cocodrilo")
	fmt.Println(aAnimals)
	fmt.Println(sPets)
	fmt.Println("Cantidad de elementos: ", len(sPets))
	fmt.Println("Tamaño total del slice: ", cap(sPets))

	sPets = append(sPets, "iguana") //* Si le vuelvo a apregar otro, lo agrega al final
	fmt.Println(aAnimals)
	fmt.Println(sPets)
	fmt.Println("Cantidad de elementos: ", len(sPets))
	fmt.Println("Tamaño total del slice: ", cap(sPets))

	//!Qué pasa si intentamos agregar un elemento más al slice, siendo que ya
	//!alcansamos el límite del Array
	sPets = append(sPets, "llama")
	sPets = append(sPets, "vaca")
	//* Si le vuelvo a apregar otro, duplica el tamaño del Slice original
	//*En este caso se le asignaron de a uno, pero si la capacidad del array
	//*se viera superada de una sola vez, crea un array nuevo con el doble de tamaño
	//*al que hace referencia y aparte devuelve los valores originalmente inicializados
	//*En el primer array
	fmt.Println(aAnimals)
	fmt.Println("Tamaño total del Array: ", len(aAnimals))
	fmt.Println(sPets)
	fmt.Println("Cantidad de elementos: ", len(sPets))
	fmt.Println("Tamaño total del slice: ", cap(sPets))

	//!Tambien se puede crear un slice sin hacer referencia a ningún Array
	sPets2 := []string{"1", "2", "3", "4"}
	fmt.Println(sPets2)
	fmt.Println("Cantidad de elementos: ", len(sPets2))
	fmt.Println("Tamaño total del slice: ", cap(sPets2))

	sPets2 = append(sPets2, "5")
	sPets2 = append(sPets2, "6", "7", "8")
	fmt.Println(sPets2)
	fmt.Println("Cantidad de elementos: ", len(sPets2))
	fmt.Println("Tamaño total del slice: ", cap(sPets2))

	sPets2 = append(sPets2, "9")
	fmt.Println(sPets2)
	fmt.Println("Cantidad de elementos: ", len(sPets2))
	fmt.Println("Tamaño total del slice: ", cap(sPets2))

	fmt.Printf("Tipo aAnimals: %T\n", aAnimals)
	fmt.Printf("Tipo sPets: %T\n", sPets)
	fmt.Printf("Tipo sPets2: %T\n", sPets2)

	//!Otra forma de declarar los slice es con una funcion make
	//*A esta funcion se le pasa por parámetros:
	//* make(Type, size, int)
	//*tipo de dato a almacenar, tamaño, y capacidad
	sProof := make([]string, 0, 3)
	fmt.Println(sProof)
	fmt.Printf("Tipo sProof: %T\n", sProof)
	fmt.Println("Cantidad de elementos: ", len(sProof))
	fmt.Println("Tamaño total del slice: ", cap(sProof))

	sProof = append(sProof, "1", "2", "3")
	fmt.Println(sProof)
	fmt.Println("Cantidad de elementos: ", len(sProof))
	fmt.Println("Tamaño total del slice: ", cap(sProof))

	//*Si nos pasamos de su capacidad, la duplica
	sProof = append(sProof, "4")
	fmt.Println(sProof)
	fmt.Println("Cantidad de elementos: ", len(sProof))
	fmt.Println("Tamaño total del slice: ", cap(sProof))

	//!Slice declarado de diferentes formas
	//var sValue []string
	//sValue := []string{}
	sValue := make([]string, 0, 0)
	fmt.Printf("Valor por defecto: %v\n", sValue)
	fmt.Println(sValue == nil)
	fmt.Println("Cantidad de elementos: ", len(sValue))
	fmt.Println("Tamaño total del slice: ", cap(sValue))
	//*nil hace referencia a null o empty
	//*El resultado va a cambiar dependiendo
	//*Si el slice esta solo declarado, nil = true
	//*Si ya esta declarado:
	//? sValue := []string{}
	//*nil = false

	sValue = append(sValue, "")
	fmt.Printf("Valor por defecto: %v\n", sValue[0])

	sValue = append(sValue, "Hola")
	fmt.Printf("Valor por defecto: %v\n", sValue[1])

	fmt.Println(sValue)
	fmt.Println("Cantidad de elementos: ", len(sValue))
	fmt.Println("Tamaño total del slice: ", cap(sValue))

	sValue = append(sValue, "Chau")
	fmt.Printf("Type: %T\n", sValue[2])

	fmt.Println(sValue)
	fmt.Println("Cantidad de elementos: ", len(sValue))
	fmt.Println("Tamaño total del slice: ", cap(sValue))
}
