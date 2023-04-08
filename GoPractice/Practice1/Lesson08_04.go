package main

import (
	"fmt"
)

func main() {
	a := [2][2]string{{"sato", "suzuki"}, {"takahashi", "onodera"}}
	fmt.Println(a[0])
	fmt.Println(a[1][1])
}
