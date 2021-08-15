package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(reflect.TypeOf("Hello World"))
	fmt.Println(reflect.TypeOf(23))
	fmt.Println(reflect.TypeOf(3.14))
	fmt.Println(reflect.TypeOf(true))
}

// 출력결과물
// string
// int
// float63
// bool
