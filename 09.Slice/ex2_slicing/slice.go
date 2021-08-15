package main

import "fmt"

func main() {
	s := make([]int, 3)
	s[0] = 1
	s[1] = 2
	s[2] = 3
	t := s[0:3]
	u := s[:3]
	v := s[1:]
	fmt.Printf("t = %d, u = %d, v = %d", t, u, v)
}

// 출력 결과물
// t = [1 2 3], u = [1 2 3], v = [2 3]
