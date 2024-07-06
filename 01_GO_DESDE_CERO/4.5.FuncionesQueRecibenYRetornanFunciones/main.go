package main

import "fmt"

func main() {
	//! FUNCIÓN QUE RECIBE FUNCIÓN
	silceNumbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(silceNumbers)

	//* Para mandarle la func por parámetro simplemente la creo, no debe tener una lógica compleja
	newSilec := filter(silceNumbers, func(num int) bool { return num >= 5 })
	fmt.Println(newSilec)

	//* También podríamos crear la función por fuera y pasarsela por parámetro
	//* Se pasa sin el paréntesis
	//* Para cuando las funciones son más complejas
	newSilec2 := filter(silceNumbers, aux)
	fmt.Println(newSilec2)

	//! FUNCIÓN QUE RETORNA FUNCIÓN
	//* Al invocar la función le pasamos en paréntesis el primer valor y seguido de otros paréntesis el segundo valor
	result := sum2(6)(8)
	fmt.Println(result)

	//* También se puede guardar la primera parte en una variable, que contendría el retorno de la segunda fucnión
	//* Y solo habrría que agregarle el segundo parámetro
	aux := sum2(10)
	aux1 := aux(8)
	fmt.Println(aux1)
	fmt.Println(aux(9))
	fmt.Println(aux(5))
	fmt.Println(aux(3))

}

//! FUNCIÓN QUE RECIBE FUNCIÓN
//* Las funciones también son tipo de datos y se pueden pasar por parámetros
//* func a(namFunc func(parameters) typeReturn) typeReturn{
//* }

// * filter filtra los números mayores a 5
func filter(sliceNumbers []int, callback func(int) bool) []int {
	//* en el make le damos el tipo de dato, tamaño (cero porque no sabemos cuantos numeros va a agarrar),
	//* limite(en este caso le ponemos como límite el tamaño del que ingresa por parámetro)
	result := make([]int, 0, len(sliceNumbers))
	for _, v := range sliceNumbers {
		if callback(v) {
			//* Se validará v >= 5 dentro de la función callbak que retorna un bool
			result = append(result, v)
		}
	}
	return result
}

func aux(num int) bool {
	return num <= 5
}

//! FUNCIÓN QUE RETORNA FUNCIÓN
//* Entendemos que la función sum2 recibe un entero y devuelve una función que a su ez recibe un entero y devuelve un entero
//* Dentro del cuerpo de la fucnión, en el return, se desarrolla la otra función
//? No entendí para que mierda puede servir esto
func sum2(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}
