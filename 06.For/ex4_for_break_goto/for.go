package main

import "fmt"

func main() {
	a := 0
	for ; a < 10; a++ {
		if a == 3 {
			continue
		}
		fmt.Println(a)
		if a == 7 {
			break
		}

		if a == 4 {
			goto L1 // a == 4가 되면 L1으로 점프
		}
	}
	fmt.Println("a의 값은 :", a)
L1:
	fmt.Println("goto문으로 도달")
}

// 출력결과물
// 0
// 1
// 2
// 4
// goto문으로 도달
