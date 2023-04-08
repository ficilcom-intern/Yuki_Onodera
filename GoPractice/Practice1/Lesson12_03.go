package main

import (
	"fmt"
)

func main() {
	func(Greeting string) {
		fmt.Println(Greeting)
	}("Hello")
}
