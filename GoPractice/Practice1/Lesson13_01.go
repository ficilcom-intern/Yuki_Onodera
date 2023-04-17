package main

import (
	"fmt"
)

type student struct{
	name string
	math, english float64
}

func main()  {
	var s student
	s.name = "sato"
	s.math = 80
	s.english = 70
	
	fmt.Println(s)
}
