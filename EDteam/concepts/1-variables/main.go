package main

import "fmt"

//Todas las variables declaradas siempre deben ser usadas
//El codigo siempre debe estar dentro de la funcion main

func main() {

	var apple string = "apple"
	var banana string = "banana"
	var naranja string = "orange"

	// apple = ":apple"
	// banana = ":banana"
	// naranja = ":orange"

	fmt.Println(apple)
	fmt.Println(banana)
	fmt.Println(naranja)
	fmt.Println(apple, banana, naranja)

	//Agrupacion de variables
	var (
		pepino string = "pepino"
		sandia string = "sandia"
		mango  string = "mango"
	)
	fmt.Println(pepino, sandia, mango)

	//Declaracion y Asignacion en una sola linea
	var lechuga, brocoli, calabaza string = "lechuga", "brocoli", "calabaza"
	fmt.Println(lechuga, brocoli, calabaza)

	//Asignacion de variable corto con :=
	jitomate, pepino, zanahoria := "jitomate", "pepino", "zanahoria"
	fmt.Println(jitomate, pepino, zanahoria)

	//Asignacion de variable dinamica
	//Una vez sea asignado el tipo de dato no se va a poder modificar el tipo de dato
	//El tipo de dato no se puede cambiar pero si el valor
	//El asignador de variable corto no se puede usar para cambiar el valor uno a uno
	numero, palomita, sillon := 1, "palomita", "sillon"
	fmt.Println(numero, palomita, sillon)

	numero = 2
	fmt.Println(numero, palomita, sillon)

	palomita, limon := "maiz", "limon"
	fmt.Println(palomita, limon)

}
