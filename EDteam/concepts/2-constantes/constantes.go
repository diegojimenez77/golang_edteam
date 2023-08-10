// package main es el paquete principal
// den donde se ejecuta nuestra aplicacion
// de EDteam

package main

import "fmt"

func main() {
	//No es posible usar el operador de variable corta para asignar constantes
	const os, domain string = "linux", "ed.team"

	//La asignacion de variable dinamica es posible
	const fruta, numero = "platano", 1

	fmt.Println(os, domain)
	fmt.Println(fruta, numero)

	const (
		so      = "Windows"
		dominio = "de.team"
	)
	fmt.Println(so, dominio)

	//iota nos ayuda a asignar valores sequenciales
	//en el siguiente ejemplo iota nos ayda a darles valor constante a los meses a partir de Feb
	//las constantes no es necesario usarlas para poder compilar
	const (
		Jan = iota + 1
		Feb
		Mar
		Abr
		May
		Jun
	)
	fmt.Println(Jan, Feb, Mar, Abr, May, Jun)

	//comentario de una sola linea
	/*comentario multilinea
	es igual que en javascript
	los comentarios de bloque se usa para las funciones
	*/

	// Name es el nombre del profesor de go
	var Name = "alejandro"
	fmt.Println(Name)
}
