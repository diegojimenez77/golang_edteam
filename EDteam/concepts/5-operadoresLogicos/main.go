package main

import "fmt"

func main() {
	// Operadores Comparacion : >, <, ==, !=, >= ,<=
	// Permiten comparar dos expresiones y dan como resultado un boolean
	fmt.Println(4 > 6)       // false
	fmt.Println(4 < 6)       // true
	fmt.Println((4 + 5) < 6) //false

	fmt.Println(4 == 4) // true
	fmt.Println(4 != 4) //false

	fmt.Println(6 >= 4) //true

	//Operadores Logicos &&, ||
	//Agrupar operaciones para producir respuesta booleana
	var age uint = 17
	fmt.Println("Es Adulto?: ", age >= 18 && age <= 60)
	fmt.Println("Es Ninio o Anciano?: ", age <= 18 || age >= 60)

	//Opreador logico Unario: ! o de negacion
	// negar una expresion booleana
	fmt.Println(!(4 == 4)) //false
	fmt.Println(4 != 4)    // false

}
