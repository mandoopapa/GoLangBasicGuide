package main

import "fmt"

func main() {
	for a := 0; a < 10; a++ {
		if a == 3 {
			continue // a == 3이 되면 continue로 다시 for문의 시작점으로
		}
		fmt.Println(a)
		if a == 4 {
			break // a == 4가 되면 break로 for 루프를 탈출
		}
	}
}

// 출력결과물
// 0
// 1
// 2
// 4
