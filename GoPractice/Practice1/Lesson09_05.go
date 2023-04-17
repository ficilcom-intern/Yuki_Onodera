package main

import (
	"fmt"
)

func main() {
	x := 10
	y := 2
	y += 10
	y --
	fmt.Println(x > y)
	fmt.Println(y)
}
