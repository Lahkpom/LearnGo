package main

import (
	"errors"
	"fmt"
	"strconv"
)

var (
	errConvert  = errors.New("Hubo un error en la converción del string a entero.")
	errNotKey   = errors.New("No existe la key")
	errConvert2 = errors.New("strconv.Atoi:")
	errConvert3 = errors.New("invalid syntax")
)

var mapFood = map[int]string{
	1: "🍔",
	2: "🍕",
	3: "🍟",
}

func main() {
	//* Go trabaja con errores, no con excepciones
	//* Hay una función string converter que permite convertir string a otros tipos de datos
	//* se llama strconv, con su función Atoi lo convertimos a entero
	//* Pero tiran error si el string qu les pasamos no contiene solo números

	//* Acá num es de tipo entero y err es de tipo error
	num, err := strconv.Atoi("10")

	fmt.Println(num)
	fmt.Println(err) //* En este caso es nil porque no hubo error

	//* Si le pasamos un string con otros caracteres:
	num, err = strconv.Atoi("asd")

	//? En este caso, num va a mantener el valor cero, y err va a contener un mensaje de error
	fmt.Println(num)
	fmt.Println(err) //* strconv.Atoi: parsing "asd",invalid syntax
	//* Pero no se corta el programa

	//*Cómo controlarlo?
	/*
		if err != nil {
			fmt.Println("Hay error")
			//*Puso la palabra return como para cortar el programa
			//? Esto hace que termine la ejecución acá y no continúa con el resto
			return
		} //*Con este if con return controlamos el flujo
	*/

	func() {
		fmt.Println("Pasé el return")
	}()

	a1, a2 := search("")
	fmt.Println(a1)
	fmt.Println(a2)
	a1, a2 = search("Hola")
	fmt.Println(a1)
	fmt.Println(a2)

	p1, p2 := search("a")
	if p2 != nil {
		fmt.Println("Error")
		fmt.Println(p2)
		return
	}
	fmt.Println(p1, ". Esto se imprime solo si no salta error")
	fmt.Println("Esto esta siendo impreso desde cambios realizados en el Project IDX")

	//! Ejemplos encadenando errores
	aux, err := search2("1")
	//* Si hay algún error lo ataja este if
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(aux)

	//*Acá tiene que saltar error por no convertir
	aux2, err2 := search2("a")
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(aux2)

	//* Acá tiene que saltar error por no existir la key
	aux3, err3 := search2("5")
	if err3 != nil {
		fmt.Println(err3)
	}
	fmt.Println(aux3)

	//! Con esta forma que ve qué función dentro de qué función fue la que causó el error

	//! Cómo comprobar la función con un error específico?
	//* En lugar de pasar el error encadenado con %v, lo pasamos con %w, que envuelve el error (Es como que en el fondo mantiene el mismo error)
	//* Y podemos chequearlo con la función error.Is(err, typeError)
	z, y := search2("a")
	fmt.Printf("%q", z)
	fmt.Println(y)
	if errors.Is(y, errConvert) {
		fmt.Println("Controlamos el error con un error.Is()")
	}

	//! NO PUDE CONTROLAR LA FUNCIÓN DE ARRIBA CON ERROR ESPECÍFICO

}

// ! CÓMO CREAR NUESTROS PROPIOS ERRORES
// ? Tenemos dos formas, con el package errors con la función error del package format
// * se crean a nivel de paquete
// * Se crean con variables a las que les ponemos a cada una el tipo de error que nos puede llegar a surgir
// * Entre comillas le pasamos el mensaje de error
var errStrEmpty = errors.New("Error - this str is empty")

// * Estos los ponemos en las funciones, como valor a devolver
func search(str string) (r string, e error) {
	//* Desarrollamos la lógica
	r = str
	if len(str) == 0 {
		e = errStrEmpty
	}
	return
}

//! MODIFICACIÓN HECHA DESDE IDX

// * Podemos ir encadenando errores para que salte uno u otro y poder determinar de dónde proviene el error
// * Tenemos precargado un map de mapFood
// * Esta función recibe un número como str y después llama a otra función que busca dentro del mapa
func search2(key string) (str string, err error) {
	//* La key es un número que se recibe como string
	num, err := strconv.Atoi(key)
	//* Si hay algún error en la conversión, agarramos ese error con este if
	if err != nil {
		//* En caso que no se pueda lograr la converisón, se retorna un str vacío y el error
		//* Se le pasa al error este el txt que le corresponde.
		// err = errConvert
		//! Otra forma de crear errores es con la función fmt.Errorf("") que imprime un tipo error
		//? Es biuen práctica poner qé función es la que está dando el error, y luego el mensaje, para identificar la función que tira el error
		return str, fmt.Errorf("search2(): strconv.Atoi(): %w", errConvert)
	}
	//* Si no salta el error anterior vamos a llamar a la función que va a hacer la búsqueda
	str, err = findFood(num)
	if err != nil {
		return str, fmt.Errorf("search2(): %w", err)
	}
	return
}

func findFood(num int) (string, error) {
	//* No hay que usar un for ya que con solo agregar la key como índice el mismo mapa nos dice
	//* Si esa key existe o no nos lo devuelve el map con un bool
	str, ok := mapFood[num]
	if !ok {
		//* Este error se le manda como parámetro a search2 y ese la manda como parámetro al main que la termina imprimiendo
		return str, fmt.Errorf("findFood(): %w", errNotKey)
	}
	//* Si el ok es true, se devuelve nil como error
	return str, nil
}
