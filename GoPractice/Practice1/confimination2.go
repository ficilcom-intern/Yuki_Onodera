package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
	i := 0
	for i <= 10 {
		fmt.Println(i)
		i++
	}
	i = 0
	for {
		fmt.Println((i))
		if i >= 10 {
			break
		}
		i++
	}
}
