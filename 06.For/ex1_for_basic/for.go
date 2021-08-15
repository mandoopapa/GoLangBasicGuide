package main

import "fmt"

func main() {
	sum := 0
	for a := 1; a <= 10; a++ {
		sum += a
	}
	fmt.Println(sum)
}

// 출력결과물
// 55
