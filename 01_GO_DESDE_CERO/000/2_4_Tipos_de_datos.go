package main

import "fmt"

func main() {
	//*TRES TIPOS DE DATOS
	//*string
	//?No se puede utilizar comilla simple, deben ser dobles
	var str string = "Hola"
	fmt.Println(str)

	//*bool
	var verdadero bool = true
	var falso bool = false

	fmt.Println(verdadero, falso)

	//*numeric
	//?Este se subdivide en otros tipos
	//?nint
	//uint8  0 a 255
	var edad uint8 = 24
	//uint16 0 a 65.535
	//uint32 0 a 4.294.967.295
	var dni uint32 = 41622287
	//uint64 0 a 18.446.744.073.709.551.615

	fmt.Println(edad, dni)
	fmt.Printf("Tipo: %T, Valor: %v\n", edad, edad)
	fmt.Printf("Tipo: %T, Valor: %v\n", dni, dni)

	//x := (dni - edad)

	//?int
	//int8  -128 a 127
	//int16 -32.768 a 32.767
	var o int16 = -3214
	//int32 -2.147.783.678 a 2.147.483.647
	//int64 -9.223.372.036.854.775.808 a 9.223.372.036.854.775.807

	//?byte = unit8

	//?rune = int32
	//Represente a un codigo unico unicode que se puede consultar
	//por ejemplo el unicode a es igual a 97
	//se inicializa poniendo el unicode en comillas simples
	var x rune = 'a'

	fmt.Printf("Tipo: %T, Valor: %v\n", x, x)

	//?float32 y float64 con decimales
	var deimal float32 = 3.14
	var deimal2 float64 = 200.5

	fmt.Printf("Tipo: %T, Valor: %v\n", deimal, deimal)
	fmt.Printf("Tipo: %T, Valor: %v\n", deimal2, deimal2)

	//!Para poder hacer operaciones entre los distintos tipos de datos es necesario hacer un casteo
	z := dni + uint32(edad)
	fmt.Println("z es igual a: ", z)

	j := o - int16(edad)
	fmt.Println(j)

	var bd float32 = 10.0 / 3
	fmt.Println(bd)
	fmt.Printf("Tipo: %T, Valor: %v\n", bd, bd)

	var b float32 = 5.0 / 2.0
	fmt.Println(b)
	fmt.Printf("Tipo: %T, Valor: %v\n", b, b)

	//!Valores por defecto
	//El valor por defecto de los string es vacío ""
	//El valor por defecto de los numeric es cero 0
	//El valor por defecto de los bool es false (no es necesario inicializarlo como false ya que se le asigna por defecto)

	//!funcion de fmt Printf()
	//*Mediante esta función podemos, utilizando unos comandos indicados en la documentación
	//*de go, pedir por ejemplo el tipo de dato y el valor de una variable
	//*Primero entre " " colocamos el comando que coresponde, luego colocamos la
	//*variable de la cual queremos esos datos, la cantidad de veces segun info pedimos
	//* %T pide el tipo de dato, %v pide el valor
	//* fmt.Printf("Tipo: %T, Valor: %v", variable, variable)
	fmt.Printf("Tipo %T, Valor %v\n", verdadero, verdadero)
	fmt.Printf("Verdadero: %v\n", verdadero)
	fmt.Printf("Tipo %T, Valor %v\n", str, str)

}
