package main

import "fmt"

func main() {
	//Operadores Aritmenticos (), *, /, %, +, - , el codigo realiza agrupaciones de acuerdo a la prioridad aritmetica
	var a = (2 + 3) * 2
	fmt.Println(a)

	//Operadores de asignacion: = , += , -=, *=, /=, %=
	var b int = 5
	b = b + 2
	fmt.Println(b)

	var c int = 5
	c += 2
	fmt.Println(c)

	var d int = 5
	d -= 2
	fmt.Println(d)

	//declaracion post-incremento y post-decremento: ++, --, no existe la declaracion preincrepento o pre-decremento
	// una expresion es una parte de codigo que produce un valor por ejemplo una operacion como las de arriba
	// una declaracion es una instruccion del lenguaje que realiza un valor, por ejemplo incrementar el valor de la variable en 1 o decrementar en 1
	var e int = 6
	e++
	fmt.Println(e)

	var f int = 8
	f--
	fmt.Println(f)
	//(no son una expresion sino un adeclaracion)
}
