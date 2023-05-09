package main

import (
	"fmt"
)

type Sasaki struct {
	Name string
	Age int
}

func main() {
	sasaki := Sasaki{
		Name: "Sasaki",
	}
	fmt.Println(sasaki)
}
