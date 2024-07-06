package main

import "fmt"

func main() {
	//* Panic hace que se detenga la ejecución del programa
	//* Ejemplo con una división, si el divisor es 0 entraremos en pánico
	division(3, 1)
	division(3, 6)
	division(3, 0)
	division(3, 3)

}

func division(a, b int) {
	//* El recover lo que hace es rescuperarse de un panic, sería algo como el try catch de las excepciones
	//* Evita que el programa se corte
	//* Para implementarlo lo combinamos con un defer y una función anónima
	//* De esta forma se ejecuta al final de la función y evalua si hay o no un panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Se ha producido un panic que se salvó por el recover")
			fmt.Println("division():", r) //* Esto muestra el mensaje del panic
		}
	}()
	validateZero(b)
	fmt.Println("Resultado: ", float32(a)/float32(b))
}

func validateZero(b int) {
	if b == 0 {
		//* El panic cortará la ejecución del programa y nos mostrará el mensaje que le indiquemos
		//* y seguido nos indicará de donde provino el alerta del panic
		//* Eso hay que leerlo de abajo a arriba para seguir el camino del panic
		panic("validateZero(): No se puede dividir entre 0!!! 🤕")
	}
}
