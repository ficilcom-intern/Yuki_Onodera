package main

import (
	"fmt"
)

type Student struct {
	name          string
	math, english float64
}

func main() {
	s := Student{"sato", 80, 70}
	fmt.Println(s)
}
