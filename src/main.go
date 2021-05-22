package main

import "fmt"

func main() {

	fmt.Println("Dig por aca")

	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("pin", pi)

	fmt.Println("pi2", pi2)

	//Declarate Variable

	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	//Zero values

	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// square Area

	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * 2

	fmt.Println("Area cuadrado:", areaCuadrado)
}
