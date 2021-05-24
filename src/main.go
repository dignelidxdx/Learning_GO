package main

import (
	"fmt"

	pk "example.com/goproject/src/mypackage"
)

func main() {
	var myCar pk.CarPublic
	myCar.Brand = "Ferrari"
	myCar.Year = 2020
	fmt.Println(myCar)
}
