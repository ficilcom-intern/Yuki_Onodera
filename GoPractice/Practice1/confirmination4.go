package main

import (
	"fmt"
)

type User struct {
	gender string
	age    float64
}

func main() {
	var s User
	s.gender = "male"
	s.age = 20
	fmt.Println(s)

	u := User{"female", 21}
	fmt.Println(u)
}
