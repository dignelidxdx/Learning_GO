package main

import (
	"fmt"

	pk "example.com/goproject/src/challenge"
)

func main() {
	myPc := pk.New(12, 200, "HP")
	myPc.SetRam(16)
	myPc.FormatPrint()
	fmt.Println("Se duplica la ram")
	myPc.DuplicateRAM()
	myPc.FormatPrint()
}
