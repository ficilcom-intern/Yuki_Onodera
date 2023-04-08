package main

import (
	"fmt"
)

func main() {
	arr := [...]int{2, 5, 6, 9}
	sum := 0
	for i := 0; i <= 3; i ++ {
		sum += arr[i]
		fmt.Println(sum)
	}
}
