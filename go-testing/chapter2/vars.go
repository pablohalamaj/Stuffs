package main

import "fmt"

var desc string
var radius int32 // Aca declaro la variable radius de manera explicita pero sin valor, se inicializa con valor default
var active bool
var satellites []string

func main() {
	var name string = "Sun" // Aca declaro la variable name como string y le asigno valor
	desc = "Star"
	radius = 685800
	mass := 1.989E+30 // Acá declaro la variable mass de manera implicita porque puse el := , por lo cual no hay que poner var y el tipo lo toma del valor designado
	active = true
	satellites = []string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}
	fmt.Println(name)
	fmt.Println(desc)
	fmt.Println("Radius (km)", radius)
	fmt.Println("Mass (kg)", mass)
	fmt.Println("Satellites", satellites)
}
