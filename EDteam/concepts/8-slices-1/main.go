package main

import "fmt"

func main() {
	//slice: son apuntadores a un Arraay, no poseen datos.
	things := [7]string{"carro", "camioneta", "camion", "autobus", "motocicleta", "trike", "bicicleta"}
	cars := things[0:4]      //go puede inferir la pocicion inicial en este caso puede usarse [:4]
	bicicleta := things[3:7] //a la posicion final se le suma 1, en este caso imprime de autobus a bicicleta
	//go puede inferir la posicion final en este caso podria colocarse [3:]
	bicicleta[1] = "ambulancia" // hacemos el cambio en el slice bicicleta en la posicion 1, cambiando motocicleta por ambulancia. pero lo cambia en todos los arrays.
	all := things[:]            // go infiere la posicion inicial y final por lo que considera todo el arrego things.

	fmt.Println("Things:", things)
	fmt.Println("Cars:", cars)
	fmt.Println("Bicicleta:", bicicleta)

	fmt.Println("Cars[0]: ", cars[0])
	fmt.Println("Bicicleta[0]: ", bicicleta[0])
	fmt.Println("ArrayCompleto: ", all)
}
