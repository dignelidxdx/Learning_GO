package main

import "fmt"

func esPar(n int) bool {
	if n%2 == 0 {
		fmt.Printf("%d es un número par.\n", n)
		return true
	} else {
		fmt.Printf("%d es un número impar.\n", n)
		return false
	}
}

func validUser(usr, pwd string) bool {
	const usrBase = "dignelidxdx"
	const pwdBase = "qwerty"
	if usr == usrBase && pwd == pwdBase {
		fmt.Println("Usuario y contraseña validos")
		fmt.Printf("¡Bienvenido %s!\n", usr)
		return true
	} else {
		fmt.Println("Usuario o contraseña incorrectos")
		fmt.Println("¡Acceso denegado!")
		return false
	}
}

func main() {
	esPar(15) //15 es un número impar.
	esPar(12) //12 es un número par.
	fmt.Println("")

	validUser("dignelidxdx", "qwerty")
	fmt.Println("")
	// Usuario o contraseña incorrectos
	// ¡Acceso denegado!

	validUser("Digneli", "12345")
	fmt.Println("")
	// Usuario o contraseña incorrectos
	// ¡Acceso denegado!

	validUser("Digneli", "qwerty")
	fmt.Println("")
	// Usuario y contraseña validos
	// ¡Bienvenido Digneli!
}
