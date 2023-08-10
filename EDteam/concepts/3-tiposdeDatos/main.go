package main

import "fmt"

func main() {
	//bool, string, numeric
	var a bool = true
	fmt.Printf("Tipo: %T, Valor: %v\n", a, a)

	//string
	var b string = "Edteam"
	fmt.Printf("Tipo: %T, Valor: %v\n", b, b)

	var c uint8 = 33
	fmt.Printf("Tipo: %T, Valor: %v\n", c, c)

	var d byte = 12
	fmt.Printf("Tipo: %T, Valor: %v\n", d, d)

	var e rune = 'e'
	fmt.Printf("Tipo: %T, Valor: %v\n", e, e)

	var f float32 = 123.24
	fmt.Printf("Tipo: %T, Valor: %v\n", f, f)
}

//Restricciones, no es posible hacer operaciones entre diferentes tipos de datos
//Go es estaticamente tipado por lo que el tipo de dato no puede cambiar
//Hay que tener cuidado con las operaciones que se realicen para no desbordar la memoria
//El identificador blank es un _ que se usa para haver validaciones se usa como si fuera una variable para los casos en los que se usan diferentes tipos de datos en una operacion.
//Valor cero, hay un valor cero para cada tipo de dato, el tipo de dato string si no se le ha sido asignado un valor se toma como un valor cero o valor vacio.
// a un valor uint0 se le asigna por default el valor cero
// para los valores bool si no han sido asignados tienen por default el valor false.
