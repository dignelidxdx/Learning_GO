package main

import (
	"fmt"
	"strings"
)

func isPalindromo(text string) {
	var textoReverse string
	for i := len(text) - 1; i >= 0; i-- {
		textoReverse += string(text[i])
	}

	if strings.ToLower(text) == strings.ToLower(textoReverse) {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es palindromo")
	}

}

func main() {
	slice := []string{"Hola", "que", "hace"}

	// EN CASO QUIERA EL CONTENIDO SIN INDICE
	for _, valor := range slice {
		fmt.Println("valor:", valor)
	}

	// EN CASO QUIERA SOLO INDICE
	for i, _ := range slice {
		fmt.Println("indice:", i)
	}

	// EN CASO QUIERA SOLO INDICE Y CONTENIDO
	for i, valor := range slice {
		fmt.Println("indice:", i, "valor:", valor)
	}

	// ama
	// amor a roma
	isPalindromo("Ama")
}
