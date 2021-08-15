package main

import "fmt"

func main() {
	i := 0
L1:
	for {
		for {
			if i == 0 {
				break L1 // 위의 L1으로 이동하는 레이블을 작성
			}
		}
	}
	fmt.Println("무사히 탈출!")
}

// 출력결과물
// 무사히 탈출!
