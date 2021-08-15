package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}
	for _, num := range numbers {
		fmt.Println(num)
	}
}

// 출력결과물
//	1
//	2
//	3
