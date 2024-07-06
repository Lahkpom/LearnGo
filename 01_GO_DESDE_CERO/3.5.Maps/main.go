package main

import (
	"fmt"
)

func main() {
	//* Los mapas son estruturas de clave valor
	//* las inicializamos con make(map[keyType]dataType)
	mMusic := make(map[string]string)

	mMusic["Guitar"] = "🎸"
	mMusic["Violin"] = "🎻"

	fmt.Println(mMusic)
	fmt.Printf("Type = %T, Value = %v, Value2 = %s, Value3 = %q, Dir = %v\n", mMusic, mMusic, mMusic, mMusic, &mMusic)

	//* Hecho de esta forma, siempre debe terminar con una coma
	mTech := map[string]string{
		"Computer": "💻",
		"Mouse":    "🖱️",
		"Keyboard": "⌨️",
	}

	mTech["Computer"] = "mmm" //* Si volvemos a repetir una clave que ya existe, vamos a reemplazar su valor, no se agrega una nueva

	fmt.Println(mTech)
	fmt.Printf("Type = %T, Value = %v, Value2 = %s, Value3 = %q, Dir = %v\n", mTech, mTech, mTech, mTech, &mTech)

	//! Para eliminar elementos
	//* delete(map, "key")
	delete(mTech, "Mouse")
	fmt.Println(mTech)

	//* Si intentamos imprimir con una llave que no existe, no da error, solo devuleve el valor cero del tipo de dato almacenado
	fmt.Printf("Type = %T, Value = %v, Value2 = %s, Value3 = %q\n", mTech["Computer"], mTech["Computer"], mTech["Computer"], mTech["Computer"])
	fmt.Printf("Type = %T, Value = %v, Value2 = %s, Value3 = %q\n", mMusic["Fake"], mMusic["Fake"], mMusic["Fake"], mMusic["Fake"])

	//* Llamar a un mapa con la llave no solo muestra el valor, sino que también devulve un bool que indica si la llave esa existe o no
	//* es decir que devuelve dos valores, primero el dato y luego el bool
	aux1, aux2 := mMusic["Fake"]
	fmt.Printf("%q, %v\n", aux1, aux2)

}
