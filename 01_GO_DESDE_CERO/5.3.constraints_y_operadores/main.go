package main

// golang.org/x/exp/constraint

func main() {

}

//* Para no tener que trabajar declarando en la T CADA tipo de dato, podemos usar un paquete constraint que tiene go que hay que importar y requiere instalación
//* Ahí nos guarda interfaces con agrupamientos para trabajar
// go get golang.org/x/exp/constraints
//* Ese de arriba es el comando para instalarlo en la pc, en project idx ya lo trae

func sum3[T int | float64](nums ...T) (sum T) {
	for _, v := range nums {
		sum += v
	}
	return
}
