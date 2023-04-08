package main

import (
	"fmt"
)

func main() {
	a := [3]string{"sato", "suzuki", "takahashi"}
	a[1] = "onodera"
	fmt.Println(a[1])
}
