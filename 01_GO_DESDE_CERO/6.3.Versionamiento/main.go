package main

import (
	//* Tenemos que importar con la url acá y aclarar qué carpeta interna quiero usar
	"strings"

	"github.com/Lahkpom/mainModules/slices"
	"rsc.io/quote"
)

//! PARA PODER USAR LOS MÓDULOS O PAQUETES QUE NOSOTROS MISMO CREAMOS, HAY QUE SUBIRLOS A UN REPOSITORIO DE GITHUB E IMPORTARLO
//! DESDE ALLÍ, PARA ESO TENEMOS QUE EJECUTAR
// go mod init github.com/usuario/proyecto
//* Se crea el archivo mod de donde go busca las cosas
//* Siempre después de agregar un modulo hay que pasar el comando
// go mod tidy

func main() {
	list := []string{"perro", "Gato", "gorila", "mono", "guacamayo", quote.Hello()}
	slices.Includes(list, "Gato")
	//fmt.Println(aux)
	//* Esto va a filtrar aquellos valores que empiecen con una letra determinada y retorna un bool de si encontró o no el valor
	slices.Filter(list, func(v string) bool {
		return strings.HasPrefix(strings.ToLower(v), "h")
	})
}
