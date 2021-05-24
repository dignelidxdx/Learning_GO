package mypackage

import "fmt"

// CarPublic Car con acceso publico
type CarPublic struct {
	Brand string
	Year  int
}

type CarPrivate struct {
	Brand string
	Year  int
}

//PrintMessage |
func PrintMessage() {
	fmt.Println("Hola")
}
