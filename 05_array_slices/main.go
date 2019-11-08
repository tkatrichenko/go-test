package main

import (
	"fmt"

	"../03_pkg/reverse"
)

func main() {
	carsArray := []string{"BMW", "Mersedes"}

	fmt.Println(carsArray)
	carsReverce := carsArray[1]
	fmt.Println(reverse.Reverse(carsReverce))
}
