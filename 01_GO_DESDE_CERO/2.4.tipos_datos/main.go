package main

import (
	"fmt"
)

func main() {
	//!bool
	var a bool = true
	var b bool = false
	//?con fmt.Printf podemos aplicar shortcuts de formato, buscar en la documentación
	//?%v imprime el valor
	//?%T imprime el tipo de dato
	//?fuera de las comillas, pondremos que variable le corresponde a cada atajo
	fmt.Printf("Tipo: %T, Valor: %v\n", a, a) //*usamos el \n para el salto de linea
	fmt.Printf("Tipo: %T, Valor: %v\n", b, b)

	//*Tener en cuenta que por defecto un bool siempre es falce

	var x bool
	var y bool = false
	fmt.Printf("Type: %T, Value: %v\n", x, x)
	fmt.Printf("Type: %T, Value: %v\n", y, y)

	var t bool = x == y
	fmt.Printf("Type: %T, Value: %v\n", t, t)

	//!string
	//*tener en cuenta que por defecto un string siempre es empty
	var c string = "Hello world"
	fmt.Printf("Tipo: %T, Valor: %v\n", c, c)
	var j string
	fmt.Printf("Tipo: %T, Valor: %v\n", j, j)

	//!numeric
	//?se subdivide en muchas variantes con sus limites

	//*unit
	//?unsigned -
	//* uint8 integers (0 to 255)
	//* uint16 integers (0 to 65535)
	//* uint32 integers (0 to 4294967295)
	//* uint64 integers (0 to 18446744073709551615)
	//* uint  depende de nuestro sistema operativo, if our os is 32bits, será uint32, y si es de 64bits será uint64

	//*int
	//?signed
	//* int8 integers (-128 to 127)
	//* int16 integers (-32768 to 32767)
	//* int32 integers (-2147483648 to 2147483647)
	//* int64 integers (-9223372036854775808 to 9223372036854775807)
	//* int  depende de nuestro sistema operativo, if our os is 32bits, será int32, y si es de 64bits será int64

	//*byte
	//?alias for uint8

	//*rune
	//?alias for int32
	//Se usa con el sistema unicode, está complicado

	//*float32

	//*float64

	var d uint8 = 25

	fmt.Printf("Type: %T, Value: %v\n", d, d)

	//! NO SE PUEDEN REALIZAR OPERACIONES ARITMÉTICAS DE FORMA DIRECTA ENTRE LOS DISTINTOS TIPOS DE NUMERIC
	//! HAY QUE CASTEAR

	var e uint16 = 2550

	f := uint16(d) + e

	fmt.Printf("Type: %T, Value: %v\n", f, f)

	//*Esto es medio raro pero permite castear bajando el rango, retransforma el número
	h := d + uint8(e)

	fmt.Printf("Type: %T, Value: %v\n", h, h)

	fmt.Println(e)
	fmt.Println(uint8(e))

	//!
	//*cuando hay una variable que tiene una logica que no queremos perder, pero al no usarla go nos da un error
	//*podemos usar un operador blank, con el underscore, que omite la línea y no da errores

	_ = uint16(d) + e //?imaginemos una lógica más complicada

	//! Valores por defecto
	var k string
	var l int
	var m float32
	var n bool

	fmt.Printf("Type: %T, Value: %v\n", k, k) //* Vacío
	fmt.Printf("Type: %T, Value: %v\n", l, l) //* 0
	fmt.Printf("Type: %T, Value: %v\n", m, m) //* 0
	fmt.Printf("Type: %T, Value: %v\n", n, n) //* false

}
