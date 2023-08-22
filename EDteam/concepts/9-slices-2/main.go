package main

import "fmt"

func main() {
	// len(): # de elementos en el slice
	// cap(): # de elementos del array origen, a partir del indice donde se creo el slice

	animals := [5]string{"gorila", "perro", "gato", "pajaro", "elefante"}
	pets := animals[1:3] // imprime perro y gato, recordemos que a la posicion final le suma 1

	// pets[2] = "caballo" // este caso marca un runtime error: index aut of range [2] with length 2 ya que pets soli tiene posicion [1:3] que es perro y gato  0,1 no tiene posicion 2

	pets = append(pets, "caballo") // modifica el arreglo original, cambia el pajaro por el caballo ya que estamos referenciando o apuntando el array original.
	pets = append(pets, "jirafa")  // vuelve a modificar el arreglo original cambiando elefante por jirafa
	pets = append(pets, "burro")   // en este caso agrega burro solo al slice pets, no al array animals ya que exede su capacidad original. la capacidad cambio a 8 ya que te da el doble.

	fmt.Println("animals: ", animals)
	fmt.Println("pets: ", pets)
	fmt.Println("tamaño pets: ", len(pets))   // imprime el numero de elementos en el slice en este caso 2.
	fmt.Println("capcidad pets: ", cap(pets)) // numero de elementos del array origen, apartir del inice donde se creo el slice perro, gato , pajaro, elefante, en este caso 4.

	pets2 := []string{"perro", "gato"}
	fmt.Println("pets2: ", pets2)
	fmt.Println("tamaño pets2: ", len(pets2))
	fmt.Println("capacidad pets2: ", cap(pets2))

	pets3 := make([]string, 0, 3)
	pets3 = append(pets3, "cerdo", "vaca", "obeja", "burro") // en este caso aumenta la capacidad a 6 ya que al agregar burro pasa la capacidad original de 3.
	fmt.Println("pets3: ", pets3)
	fmt.Println("tamaño pets3: ", len(pets3))
	fmt.Println("capacidad pets3: ", cap(pets3))

	var pets4 []string
	fmt.Println("pets4: ", pets4)
	fmt.Println("tamaño pets4: ", len(pets4))
	fmt.Println("capacidad pets4: ", cap(pets4))
	fmt.Println("valor cero: ", pets4 == nil) // muestra true ya que el slice esta en cero

	pets5 := []string{}
	fmt.Println("valor cero: ", pets5 == nil) // muestra false ya que al inicializar con las {} ya no esta vacio.
}
