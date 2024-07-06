package main

import "fmt"

func main() {
	//* Los SLICES son punteros a arrrays, por lo que no poseen valor en sí, solo referencia de dirección en memoria de un array
	aNames := [...]string{"Leo", "Fiore", "Facu", "Gonza", "Ale"}
	fmt.Println(aNames)

	//* Al Slice se lo declara asignandole un array aclarando entre corchetes desde qué indice a qué indice va a almacenar la dirección
	//* se separan por : [inicio:final]
	//* Si queremos que empiece por el cero, no hace falta ponerlo
	//* De igual forma si quiero que vaya hasta el final, podemos dejar vacío después del :
	sNames := aNames[:2] //* Al indice final no lo incluye
	sNames2 := aNames[2:4]
	fmt.Println(sNames)
	fmt.Println(sNames2)

	sNames = aNames[:(2 + 1)] //* Para hacerlo incluyente hay que sumarle 1
	fmt.Println(sNames)

	//* Lo que modificamos en el slice se modifica en el array original
	sNames2[0] = "FLAG"
	fmt.Println(aNames)
	fmt.Println(sNames)
}
