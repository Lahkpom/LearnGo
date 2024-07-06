package main

import "fmt"

//* Las estrucrturas son el equivalente a las clases
//* Se pueden almacenar diversos tipos de datos
//* type nameStruct struct{
//* 	nameVar1 dataType
//* 	nameVarN dataType
//* }

//* Se declaran fuera de la función main, a nivel de paquete, para que estén visibles por todo el archivo

type Person struct {
	name      string
	age       uint8
	isMarried bool
}

func main() {
	//* para crear instancias
	p1 := Person{
		name: "Leonel",
		age:  25,
	}

	//* Si no especificamos alguna de las variables, se setea con el valor cero
	//? El + es para que también imprima el nombre de cada campo y no solo los valores
	fmt.Printf("%+v\n", p1)

	//* Si vamos a poner los valores en el mismo orden que están en la estructura, podemo no escribir el nombre de la variable
	//* y simplemente separarlos con ,
	p2 := Person{"Facu", 23, false}

	fmt.Printf("%+v\n", p2)

	//* Para acceder a cada campo es struct.varName
	fmt.Println(p1.name)
	fmt.Println(p1.age)
	fmt.Println(p1.isMarried)

	//* También podemos hacer punteros
	aux := &p2

	fmt.Println(aux)

	//* Si modifico algo en el puntero se va a modificar también en el original

	aux.name = "facundo"
	fmt.Println(aux)
	fmt.Println(p2)
	fmt.Printf("Type: %T, Dir: %v", aux, &aux) //* Dirección en memoria de la instancia

	//* En el video lo hace
	//? (*aux).name = "value"
	//* Pero a mi me salió bien de la otra forma directa

}
