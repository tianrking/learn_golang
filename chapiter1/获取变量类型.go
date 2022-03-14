package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(reflect.TypeOf(42))
	fmt.Println(reflect.TypeOf(4.12))
	fmt.Println(reflect.TypeOf('a'))
	fmt.Println(reflect.TypeOf("ab"))
	fmt.Println(reflect.TypeOf(true))
}
