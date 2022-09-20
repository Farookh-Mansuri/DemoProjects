package main

import (
	"fmt"
)

func add(x int, y int) int {
	fmt.Println(x, y)

	for i := 1; i <= y; i++ {
		x++
	}

	return x
}

func main() {

	a := 10
	b := 27
	c := add(a, b)
	fmt.Println("Sum of ", a, " and ", b, " is ", c)
}
