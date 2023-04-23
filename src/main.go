package main

import (
	"fmt"
	"reflect"
)

func main() {
	num := 1
	str := "hello world"

	fmt.Println(num)
	fmt.Println(str)

	fmt.Println(reflect.TypeOf(num))
	fmt.Println(reflect.TypeOf(str))
}