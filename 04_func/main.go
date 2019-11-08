package main

import (
	"fmt"
)

func getSum(num1, num2 int) int {
	return num1 + num2
}

func main() {

	fmt.Println("Hello!")

	// res := math.Sqrt(b)
	fmt.Println(getSum(5, 6))
}
