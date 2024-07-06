package main

import "fmt"

func main() {
	//* For es el único ciclo de control, no hay while ni do while
	//* no va entre paréntesis

	//! For estándar
	//* declaración de inicio; condición; incremento
	count := 10
	for i := 0; i < count; i++ {
		fmt.Println(i)
	}

	//! For continuo, similar al while
	//* La declaración de inicio y el incremento no es necesario declararla en el for
	//* y pueden estar dentro del ciclo, de ese modo solo queda la condición
	//* Por lo que el for se ejecutará mientras esta sea verdadera
	i := 0
	for i < count {
		fmt.Println(i)
		i++
	}

	//! For forever que se ejecuta para siempre, indefinidamente
	//* No se pone nada más que el for
	//* Esto es útil en los servidores cuando queremos que un scrip se ejecute constantemente para algún tipo de control
	//* Por ejemplo en un websocket o para trabajar con concurrencias, que "escuche constantemente"
	i = 0
	for {
		//* En todo caso podemos ponerle un control con un break para que frene la ejecución ante alguna condición
		if i == 10 {
			break
		}
		fmt.Println(i)
		i++
	}

	//! For each
	//* for i, v := range array/slice
	//* dond i es el índice y v la variable que guarda el calor
	sliceNombres := []string{"Leo", "Facu", "Ale"}
	for i, v := range sliceNombres {
		fmt.Println("Indice: ", i, ", Valor: ", v)
	}

	//* Tambien se puede pasar el contenido del slice directo al for
	for i, v := range []string{"Leo", "Facu", "Ale"} {
		fmt.Println("Indice: ", i, ", Valor: ", v)
	}

	//* En el caso que no necesitemos el índice podemos dejarlo con un underscore _
	for _, v := range []string{"Leo", "Facu", "Ale"} {
		fmt.Println("Valor: ", v)
	}

	//* Los valores que toma v no son los reales, sino que crea un valor copia tipo sandbox, lo que hace que cualquier modificación
	//* que querramos hacerle a ese valor no lo va a impactar en el original
	//* Para poder modificar el cuerpo original tenemos que usar el i
	sliceNumbers := []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(sliceNumbers)

	//* Esto no modifica el valor en la estructura original
	for _, v := range sliceNumbers {
		v *= 2
	}
	fmt.Println(sliceNumbers)

	//* Esto sí modifica el valor de la estructura original
	for i, v := range sliceNumbers {
		sliceNumbers[i] *= v
	}
	fmt.Println(sliceNumbers)

	sliceNumbers = []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//* Si no utilizáramos la v podemos omitirla
	for i := range sliceNumbers {
		sliceNumbers[i] *= sliceNumbers[i]
	}
	fmt.Println(sliceNumbers)

	//* Ejemplo con mapas
	mapFood := map[string]string{
		"Pizza":     "🍕",
		"Hamburger": "🍔",
		"apple":     "🍎",
		"Hotdog":    "🌭",
	}

	//* En este caso automáticamente la k (i) no itera del 0 al size, sino por llaves
	for k, v := range mapFood { //* k de key y v de value
		fmt.Println("Key", k, "Value", v)
	}

	//* Podemos iterar un string
	for i, v := range "Leonel" {
		fmt.Println(i, v) //* Esto me devuelve el código de rune
	}

	for _, v := range "Leonel" {
		//* Para poder imprimir la letra tenemos que hacer un casteo a string(v)
		fmt.Println(string(v))
	}

}
