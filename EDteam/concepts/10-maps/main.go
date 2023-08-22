package main

import "fmt"

func main() {
	music := make(map[string]string)
	music["guitarra"] = "guitarra"
	music["violin"] = "violin"

	fmt.Println(music)

	tech := map[string]string{
		"computer": "computadora",
		"mouse":    "mouse",
	}

	fmt.Println(tech)

	delete(tech, "computer")
	fmt.Println(tech)

	fmt.Println(music["guitarra"])
	fmt.Println(music["fake"]) // nos retorna un elemento vacio ya que es un elemento que no existe.

	// content, ok := music["fake"] // el ok es una comprobacion regresa false o true, en este caso regresa false ya que fake no existe
	content, ok := music["guitarra"] // regrasa true ya que guitarra si existe

	fmt.Println(content, ok)
}
