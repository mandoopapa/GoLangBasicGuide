package main

import "fmt"

func main() {
	a := 1
	b := 2
	swap(a, b)
}

func swap(a int, b int) {
	temp := a
	a = b
	b = temp
	fmt.Printf("a = %d, b = %d", a, b)
}

// 출력 결과물
// a = 2, b = 1
