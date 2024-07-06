package main

import "fmt"

// golang.org/x/exp/constraint"
//*No me deja importar este modulo, pero basicamente es importar interfaces ya creadas en ese paquete, las paso
//*a crear yo mismo y las pongo acá abajo

type Complex interface {
	~complex64 | ~complex128
}

type Float interface {
	~float32 | ~float64
}

type Integer interface {
	Signed | Unsigned
}

type Ordered interface {
	Integer | Float | ~string
}

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type MyNumber int

func main() {
	var a MyNumber = 10
	var b MyNumber = 20
	fmt.Println(sum3(a, b))
	fmt.Println(isInclude([]int{1, 2, 3}, 2))
	fmt.Println(isInclude([]string{"1", "2", "3"}, "2"))
	fmt.Println(isInclude([]string{"1", "2", "3"}, "l"))
}

//* Para no tener que trabajar declarando en la T CADA tipo de dato, podemos usar un paquete constraint que tiene go que hay que importar y requiere instalación
//* Ahí nos guarda interfaces con agrupamientos para trabajar
// go get golang.org/x/exp/constraints
//* Ese de arriba es el comando para instalarlo en la pc, en project idx ya lo trae

func sum3[T Integer | Float](nums ...T) (sum T) {
	for _, v := range nums {
		sum += v
	}
	return
}

//! CONSTRAINT COMPARABLE, VIENE INCLUIDO CON GO
//* Usar restricciones tipo para operadores de comparación
//* == o !=
//* comparable es un tipo de dato que usamos para comparar
//* En este caso vamos a buscar un valor T dentro de una lista T

func isInclude[T comparable](list []T, value T) (r bool) {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return
}

// flagg
