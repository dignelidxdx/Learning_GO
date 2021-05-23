package main

import (
	"fmt"
)

func main() {

	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Printf("\n")

	//For while
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	//For forever
	counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++

		if counterForever == 50 {
			break
		}
	}

	//Countdowm
	count := 10
	for count >= 0 {
		fmt.Println(count)
		count--
	}
}
