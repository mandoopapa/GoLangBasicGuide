package main

import "fmt"

func main() {
	s := make([]int, 0, 3)
	for i := 0; i < 5; i++ {
		s = append(s, i)
		fmt.Printf("cap : %v, len : %v, add : %p\n", cap(s), len(s), s)
	}
}

// 출력결과물
// cap : 3, len : 1, add : 0xc00000e168
// cap : 3, len : 2, add : 0xc00000e168
// cap : 3, len : 3, add : 0xc00000e168
// cap : 6, len : 4, add : 0xc00000a3f0
// cap : 6, len : 5, add : 0xc00000a3f0
