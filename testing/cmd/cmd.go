package main

import (
	"fmt"
	"reflect"
)

func main() {

	tar := func(in string) string { return "Hello world" }
	functionType := reflect.TypeOf(tar)

	fmt.Println(functionType)
}
