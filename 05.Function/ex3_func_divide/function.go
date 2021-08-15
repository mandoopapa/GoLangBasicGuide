package main

import "fmt"

func main() {
	mok, suc := divide(4, 2)
	fmt.Println(mok, suc)
}
func divide(a int, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}

// 출력 결과물
// 2 true
