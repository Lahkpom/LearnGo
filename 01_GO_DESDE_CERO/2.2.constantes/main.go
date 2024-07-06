package main

import "fmt"

//*Se inicializa con const y debe tener si o si el valor de la declaración en la misma línea
//*Se recomienta usarlos a nivel paquete para que esten disponibles globalmente
const os string = "lunix"
const os2, os3 string = "windows", "macos"

//*no podemos usar el asignador orto := ya que necesitamos especificar el const
//*pero podemos omitir el tipo de dato, dejando que el dato asignado lo interprete en la aignación del valor
const proof = "proof"

const (
	proof2 = "p2"
	proof3 = "p3"
)

//*En el caso que tengamos que usar constantes incrementales, para no poner uno a uno los numeros
//* existe el operador iota, que automaticamente asigna a las variables dentro del mismo bloque
//*un numero secuencial

const (
	Jan = iota * 1 //*iota empieza en cero, puede ser * 2 la secuencia incremental
	Feb
	Mar
	Apr
	May
	Jun //*debería ser = a 6
)

func main() {

	fmt.Println(os, os2, os3, proof, proof2, proof3)
	fmt.Println(Jun)
}
