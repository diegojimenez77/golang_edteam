package main

import "fmt"

//Los punteros son variables que almacenan la direccion en memoria de un valor

func main() {
	var color string = "red"
	var pointerColor *string
	pointerColor = &color
	*pointerColor = "azul" //esta cambiando el valor original de la variable color desde el puntero

	fmt.Printf("Tipo: %T, Valor: %s\n", color, color)
	//el & nos permite obtener la direccion en memoria de una variable
	fmt.Printf("Tipo: %T, Valor: %s, Direccion: %v\n", color, color, &color)
	fmt.Printf("Tipo: %T, Valor: %vm, Desreferenciacion: %s\n", pointerColor, pointerColor, *pointerColor)
}
