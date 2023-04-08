package main

import (
	"fmt"
)

func main() {
	x := 10
	y := 10
	if age := x + y; age >= 20 {
		fmt.Println("adult", age)
	} else if age == 0 {
		fmt.Println("baby")
	} else {
		fmt.Println("child")
	}
}
