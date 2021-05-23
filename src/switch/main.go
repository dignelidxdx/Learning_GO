package main

import "fmt"

func main() {
	// Estructura condicional - Switch
	var dia int8 = 4

	switch dia {
	case 1:
		fmt.Println("Lunes")
	case 2:
		fmt.Println("Martes")
	case 3:
		fmt.Println("Miercoles")
	case 4:
		fmt.Println("Jueves")
	case 5:
		fmt.Println("Viernes")
	case 6:
		fmt.Println("Sabado")
	case 7:
		fmt.Println("Domingo")
	default:
		fmt.Println("Ese no es un día valido de la semana!")
	}

	palabra := "en platzi nunca paramos de aprender"
	a, e, i, o, u := contadorVocales(palabra)
	fmt.Printf("la frase '%s' tiene: \n", palabra)
	fmt.Printf("%d vocales a \n", a)
	fmt.Printf("%d vocales e \n", e)
	fmt.Printf("%d vocales i \n", i)
	fmt.Printf("%d vocales o \n", o)
	fmt.Printf("%d vocales u \n", u)

	queTrimestreEs(10)
}

func contadorVocales(palabra string) (int, int, int, int, int) {
	conta := 0
	conte := 0
	conti := 0
	conto := 0
	contu := 0
	for _, valor := range palabra {
		variable := string(valor)
		switch variable {
		case "a":
			conta++
		case "e":
			conte++
		case "i":
			conti++
		case "o":
			conto++
		case "u":
			contu++
		}
	}
	return conta, conte, conti, conto, contu
}

func queTrimestreEs(mes int8) {
	switch {
	case mes <= 3:
		fmt.Println("Primer Trimestre")
	case mes > 3 && mes <= 6:
		fmt.Println("Segundo Trimestre")
	case mes > 6 && mes <= 9:
		fmt.Println("Tercer Trimestre")
	case mes > 9 && mes <= 12:
		fmt.Println("Cuarto Trimestre")
	default:
		fmt.Println("Este no es un mes valido")
	}
}
