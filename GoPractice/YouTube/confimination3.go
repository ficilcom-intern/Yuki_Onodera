package main

import (
	"fmt"
)

func cal(x, y int) (int, int) {
	return x, y
}
func main() {
	x, y := cal(10, 5)
	fmt.Println(x + y)

}
