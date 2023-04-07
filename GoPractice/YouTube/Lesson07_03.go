package main

import (
	"fmt"
	"reflect"
)

func main() {
	var string_a = "Hello World"
	string_b := "Hello World"
	fmt.Println(string_a)
	// var num02 int = 1111
	// num03 := 3.14
	fmt.Println(reflect.TypeOf(string_a))
	fmt.Println(reflect.TypeOf(string_b))
}
