package main

import (
	"fmt"
	"os"
)

func main() {
	//* Defer viene de diferir
	//* Diferir sería aplazar, aplazar la ejecución de alguna función
	//* Se aplazará hasta que la función en la que el defer fu llamado retorne
	nums()
	fmt.Println("")
	numsWithDefer()
	fmt.Println("")
	nums3()

	//* Esto puede utilizarse para limpiar recurso, cerrar archivo y conexciones
	//* de red o cerrar controladores de BD

	//! Ejemplo con una creación de archivo
	createOS()
}

func nums() {
	fmt.Println("1")
	fmt.Println("2")
	fmt.Println("3")
}

func numsWithDefer() {
	//* En este ejemplo al llamar a esta función vemos que en pantalla se imprime normal 1 2 3
	//* Esto se debe a que la ejecución de aquella instrucción que tenga defer se hará a lo último
	//* Si tenemos más de un defer en la misma función, el que esté más arriba va a ser el úitmo en ejecutarse
	//* Esto es porque go va metiendo las funciones diferidas en un pila, de modo que la primera que entra es
	//* la última en salir
	defer fmt.Println("4")
	defer fmt.Println("3")
	fmt.Println("1")
	fmt.Println("2")
}

func nums3() {
	//* Acá vemos que la instrucción con defer por más que se imprima a lo último
	//* guarda el valor que tenía al momento de ejecutarse, por lo que los posteriores
	//* cambios a las variables no lo afectan
	num := 4
	defer fmt.Println(num)
	num = 5
	fmt.Println(num)
}

func createOS() {
	//* Esta func con el pckage os crea un archivo, si no lo puede crear devuelve un error
	//* Si logra crearlo, nos devuelve un puntero a ese archivo
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	//* Este comando cierra ese archivo, gracias al defer podemos ponerlo acá y dejar
	//* que se ejecute a lo último, ya que si tuvieramos que cerrar el archivo al final
	//* en caso que la siguiende instrucción write tenga un error y retorne, nunca se
	//* ejecutaría el close y quedaría el archivo abierto
	defer file.Close()

	//* Esto escribe dentro del file y nos devuelve el número de bytes que se escribieron en el archivo
	//* Usamos un blank ya que en este caso no nos interesa realmente lo que devuelve
	a, err := file.Write([]byte("Escribimos esto dentro del archivo"))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Archivo creado")
	fmt.Println("a: ", a)

	//* file.Close()

}
