package main

import (
	"fmt"
)

func main() {
	if age := 0; age >= 20 {
		fmt.Println("adult")
	} else if age == 0 {
		fmt.Println("baby")
	} else {
		fmt.Println("child")
	}
}
