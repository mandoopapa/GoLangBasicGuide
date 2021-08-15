package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	l := len(s)
	fmt.Println(l, cap(s), s)
}

// 출력결과물
// 6 6 [0 1 2 3 4 5]
