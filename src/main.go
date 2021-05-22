package main

import (
	"fmt"
	"math"
)

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

	//Data Types
	var shortInt int8 = 3
	var longInt int64 = 2313212113234256876
	var shortFloat float32 = 3.1
	var longFloat float64 = 3.153465212456432145668723312
	var text string = "string"
	var boolean bool = true
	var complex complex128 = 10 + 8i

	fmt.Println(shortInt, longInt, shortFloat, longFloat, text,
		boolean, complex)

	helloMessage := "Hello"
	worldMessage := "world"

	// Println: Salto de Linea Automatico
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	// Printf
	nombre := "La academia"
	cursos := 500
	// Con valores seguros
	fmt.Printf("%s tiene más de %d cursos\n", nombre, cursos)
	// Con valores inseguros
	fmt.Printf("%v tiene más de %v cursos\n", nombre, cursos)

	// Sprintf
	message := fmt.Sprintf("%v tiene más de %v cursos\n", nombre, cursos)
	fmt.Println(message)

	// Tipo de datos:
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T\n", cursos)

	fmt.Printf("Circulo %.2f \n", areaaCirculo(2))
	fmt.Printf("Rectangulo %.2f \n", areaaRectangulo(5, 10))
	fmt.Printf("Trapezoide %.2f \n", areaTrapezoide(10, 5, 3))

	fmt.Println("Function program initialized!")

	messsage := "First message in a function example"
	firstFunction(messsage)

	threeArguments(15, 7, "String!")

	fmt.Println(returningFunction(4))

	num, twiceNum := doubleReturn(2)
	fmt.Printf("%d multiplied by 2 is: %d\n", num, twiceNum)

	_, twiceEight := doubleReturn(8)
	fmt.Printf("8 multiplied by 2 is: %d\n", twiceEight)
}

func areaaCirculo(radio float64) float64 {
	return math.Pi * radio * radio
}
func areaaRectangulo(base float64, altura float64) float64 {
	return base * altura
}

func areaTrapezoide(B float64, b float64, h float64) float64 {
	return h * (B + b) / 2
}

func firstFunction(message string) {
	fmt.Println(message)
}

func threeArguments(first, second int, third string) {
	fmt.Printf("Two numbers: %d and %d. One string: %s\n", first, second, third)
}

func returningFunction(number int) int {
	return number * 2
}

func doubleReturn(number int) (c, d int) {
	return number, number * 2
}
