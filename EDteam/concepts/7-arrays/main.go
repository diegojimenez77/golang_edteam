package main

import "fmt"

func main() {
	//en go los arrays son de tamanio fijo, una vez declarado no se puede cambiar
	// var flags [3]string
	// flags[0] = "mex"
	// flags[1] = "black"
	// flags[2] = "white"

	// fmt.Println(flags)

	flags := [...]string{"usa", "red", "orange", "pirata"}
	fmt.Println(flags)
	//con los [...] go infiere el tamanio del array, aun asi sigue siendo un array de tamanio fijo.
}
