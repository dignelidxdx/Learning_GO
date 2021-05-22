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

	//Suma

	x := 10
	y := 50
	result := x + y
	fmt.Println("Suma", result)

	//Resta
	result = y - x
	fmt.Println("Resta", result)

	//Mutiplicacion
	result = x * y
	fmt.Println("Multiplicacion", result)

	//Division
	result = y / x
	fmt.Println("Division", result)

	//Modulo
	result = y % x
	fmt.Println("Modulo", result)

	//Incremental
	x++
	fmt.Println("Inclemental", x)

	//Decremental
	x--
	fmt.Println("Decremental", x)

	// Rectángulo
	baseRectangulo := 20
	alturaRectangulo := 10

	areaRectangulo := baseRectangulo * alturaRectangulo

	fmt.Println("El Area del Rectángulo es :", areaRectangulo)

	// Circulo : AreaCirculo = pi por radio al cudrado
	const PI float64 = 3.14 // Constant
	var radioCirculo float64 = 10

	areaCirculo := PI * radioCirculo * radioCirculo

	fmt.Println("El Area del Circulo es :", areaCirculo)

	// Trapecio
	var baseUno float64 = 6
	var baseDos float64 = 15
	var alturaTrapecio float64 = 25

	areaTrapecio := ((baseUno + baseDos) * alturaTrapecio) / 2

	fmt.Println("El Area del Trapecio es :", areaTrapecio)
}
