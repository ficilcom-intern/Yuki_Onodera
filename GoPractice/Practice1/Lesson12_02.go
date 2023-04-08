package main

import (
	"fmt"
)

func cal(x, y int) (int, int) {
	a := x + y
	b := x - y
	return a, b
}

func main() {
	result1, result2 := cal(6, 3)
	fmt.Println(result1, result2)
}
