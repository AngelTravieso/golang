package main

import (
	"fmt"
	"reflect"
)

func ReflectTypeOf() {
	var string1 string = "texto"
	double := 32.64
	isTrue := true

	fmt.Println(reflect.TypeOf(string1))
	fmt.Println(reflect.TypeOf(double))
	fmt.Println(reflect.TypeOf(isTrue))
}
