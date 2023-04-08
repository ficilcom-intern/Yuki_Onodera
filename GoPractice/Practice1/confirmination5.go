package main

import (
	"fmt"
)

type User struct {
	name string
}

func (U User) cal(weight, height float64) (result float64) {
	result = (weight / height / height) * 10000
	return
}

func main() {
	a001 := User{"sato"}
	fmt.Println(a001.name, a001.cal(80, 170))
}
