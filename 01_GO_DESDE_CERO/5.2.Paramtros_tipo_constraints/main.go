package main

import "fmt"

func main() {
	// fmt.Println(sum(2, 5, 7))
	//? AL LLAMAR A LA FUNCIÓN SE ESPECÍFICA ENTRE CORCHETES CON CUÁL TIPO DE DATO DE LOS PERMITIDOS DESEAMOS TRABAJAR
	fmt.Println(sum[int](2, 5, 7))
	fmt.Println(sum2[int](2, 5, 7))
	fmt.Println(sum2[float64](2, 5.2, 7))
	fmt.Println(div[float64](15, 2))
	fmt.Println(div[int](15, 2))          //* Acepta divisiones entre int, los trunca
	fmt.Println(sum3[MyInt](a, MyInt(b))) //* También es válido castear

	//? NO ES NECESARIO PONER LO QUE ESTÁ ENTRE CORCHETES A VECES YA QUE GO LE ASIGNA EL TYPO DE DATO DEL PRIMER ELEMENTO DEL PARÁMETRO
	//? simplemente llamamos a suma sin paréntesis
	fmt.Println(sum2(2, 5, 7))

}

/* Any no serviría para estos casos ya que no puede sumar todos los tipos de datos
func sum(nums ...any) (sum any) {
	for _, v := range nums {
		sum += v
	}
	return
}
*/

// * Para poder superar esto, existen los parámetors de tipo, que en este caso de la suma, nos permite restringir
// * los tipo de dato ingresado a aquellos que se puedan sumar
//! Hay tres tipos de constraints: Arbitrarios, de unión de elementos, de aproximación de elementos

// ? ESTE ES DE TIPO ARBITRARIO, ADMITE SOLO UN TIPO DE DATO
func sum[T int](nums ...T) (sum T) {
	for _, v := range nums {
		sum += v
	}
	return
}

// ? CONSTRAINT DE TIPO DE UNIÓN DE ELEMENTOS
// ? SE INDICA CON CORCHETES ENTRE EL NOMBRE DE LA FUNCIÓN Y LOS PARÉNTESIS
// ? DE ESTA FORMA INDICAMOS ARBITRARIAMENTE DE QUÉ TIPO DE DATO PUEDE SER T
// ? ESTO IGUAL LIMITA A QUE T SIEMPRE VA A SER DEL MISMO TIPO QUE TOME CADA VEZ
// ? SE SEPARA CADA TIPO CON UN PYPE
func sum2[T int | float64](nums ...T) (sum T) {
	for _, v := range nums {
		sum += v
	}
	return
}

func div[T int | float64](a, b T) T {
	return a / b
}

// ? TIPOS DECLARATIVOS ES CREAR YO UN TIPO DE DATO (COMO LOS QUE USAMOS EN P.O.O.
// ? SI YO CREO UN NUEVO TIPO DE DATO Y LO HAGO DE TIPO ENTERO, NO PODRÍA TRABAJAR CON LA MISMA FUNCIÓN
// ? PORQUE AHORA ES DE OTRO TIPO DE DATO
type MyInt int
type MyInt2 MyInt

var a MyInt = 5
var b MyInt2 = 10

//? a ES DE TIPOY MyInt, no int
//? PARA PODER USAR MyInt, TENGO QUE INDICARLE A LA FUNCIÓN QUE DEBE ACEPTAR ENTEROS Y CUALQUIER DERIBADO DE ENTEROS
//? se le agrega el ~ al tipo de dato para indicarle que acepta cualquier deribado

func sum3[T ~int | float64](nums ...T) (sum T) {
	for _, v := range nums {
		sum += v
	}
	return
}
