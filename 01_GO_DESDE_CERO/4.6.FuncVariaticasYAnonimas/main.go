package main

import "fmt"

func main() {
	// ! Funciones Variáticas
	a := sum(1)
	b := sum(1, 2)
	c := sum(1, 2, 3)
	d := sum(1, 2, 3, 4)
	e := sum(1, 2, 3, 4, 5)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	//* Incluso puede no recibir ningún valor y usará el valor cero del tipo de dato
	f := sum()
	fmt.Println(f)

	//? Prueba con strings
	g := sum2()
	h := sum2("Hola")
	i := sum2(h, "Como")
	j := sum2(i, "Estás", "?")
	fmt.Println(g)
	fmt.Printf("%q\n", g)
	fmt.Println(j)
	fmt.Printf("%q\n", j)

	// ! Funciones anónimas
	//* Le podemos asignar la función a una variable (variables de tipo función)
	hello := func() {
		fmt.Println("🤚")
	}
	//* Para llamarla ponemos el nombre de la variable y los paréntesis
	hello()

	//*Funciones anónimas con parámetros
	hello2 := func(a string) {
		fmt.Println(a + " 🤚")
	}
	hello2(j)

	//*Funciones anónimas con retorno
	hello3 := func(a string) (r string) {
		return a + " 🤚"
	}
	aux := hello3(j)
	fmt.Println(aux)

	//*Se le puede indica a go que autoejecute la función anónima cuando llegue a la línea donde esta esté
	//* No se la agrega a ninguna variable y se colocan () al final de las {}
	func() {
		fmt.Println("Esta función se autoejecuta")
	}()

	func() {
		//* Puedo usar las variables sin pasarlas por parámetro
		fmt.Println("Esta función se autoejecuta " + j)
		//* Acá si yo modifico j debería modificarse el real
		//* Si lo modifico en la función de abajo solo se modificaría la copia
		j = "Esta es la nueva j"
	}()

	//* Ejemplo pasándole variables por parámetro
	//* En el paréntesis de abajo se le pasa la variable a utilizar
	func(a string) {
		//* Puedo usar las variables sin pasarlas por parámetro
		fmt.Println(a)
		fmt.Println("Esta función se autoejecuta " + a)
		a = "Esta es la a"
		fmt.Println(a)
	}(j)
	fmt.Println(j)

}

// ! Funciones Variáticas
// * Permite recibir un número n de parámetros
// * Al momento de determinar el tipo de dato se le anteponen ...
// * Estos 3 puntos le indican a la función que debe convertir la variable en un slice
// * Lo que le permite recibir esta infinita cantidad de números
// ? No sería correcto poner (nums []int) porque la función esperará un slice ya creado
func sum(nums ...int) int {
	//* Debemos iterar en el slice
	var sum int
	for _, v := range nums {
		sum += v
	}
	return sum
}

func sum2(str ...string) string {
	var aux string
	for _, v := range str {
		aux += " " + v
	}
	return aux
}

//* Podemos omitir varias cosas con nombrar el retorno
func sum3(nums ...int) (sum int) {
	//* Debemos iterar en el slice
	for _, v := range nums {
		sum += v
	}
	return
}

//! Funciones anónimas son funciones sin nombre
