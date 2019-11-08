package main

import "fmt"

func main() {
	x := 1
	y := 4

	if x < y {
		fmt.Printf("%d is less then %d\n", x, y)
	}

	// if
	color := "red"
	if color == "red" {
		fmt.Println("red")
	} else {
		fmt.Println("not red")
	}

	// switch
	switch color {
	case "red":
		fmt.Println("red")
	default:
		fmt.Println("not red")

	}

}
