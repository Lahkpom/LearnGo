package main

import "fmt"

func main() {
	//*Se declara con un make(map[llave]tipoValor)
	mMusic := make(map[string]string)

	//*Para agregar valores debemos indicar la llave entre corchetes
	mMusic["guitar"] = "guitarra"
	mMusic["violin"] = "Danna"

	fmt.Println(mMusic)
	fmt.Printf("Type: %T, Value: %v\n", mMusic, mMusic)
	fmt.Println("Size: ", len(mMusic))

	//*Declararla e inicializarla
	mTech := map[string]string{
		"computer": "computadora",
		"mouse":    "ratón",
		"Keyboard": "teclado",
	}
	fmt.Println(mTech)

	mTech["computer"] = "parlante"
	fmt.Println(mTech)

	mTech["desk"] = "escritorio"
	fmt.Println(mTech)

	//*Para eliminar elementos utilizamos la función delete
	delete(mTech, "mouse")
	fmt.Println(mTech)

	fmt.Println(mTech["desk"])

}
