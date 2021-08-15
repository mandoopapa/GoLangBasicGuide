package main

import "fmt"

func main() {
	var a interface{}
	fmt.Println(a)

	a = 3
	fmt.Println(a)

	a = "Mandoo"
	fmt.Println(a)
}

// 출력 결과물
// <nil>
// 3
// Mandoo
