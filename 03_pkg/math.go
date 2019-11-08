package main

import (
	"fmt"
	"math"

	"../03_pkg/reverse"
)

func main() {
	// var a int8 = 3
	var b float64 = 5.2

	fmt.Println("Hello!")

	// res := math.Sqrt(b)
	fmt.Println(math.Sqrt(b))
	fmt.Println(reverse.Reverse("dsa"))
}
