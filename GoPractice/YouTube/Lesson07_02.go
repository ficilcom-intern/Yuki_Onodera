package main

import (
	"fmt"
	"reflect"
)

func main() {
	num01 := 123
	var num02 int = 1111
	num03 := 3.14
	fmt.Println(reflect.TypeOf(num01))
	fmt.Println(reflect.TypeOf(num02))
	fmt.Println(reflect.TypeOf(num03))
}
