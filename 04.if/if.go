package main

import "fmt"

func main() {
	var score = 0
	fmt.Scan(&score)
	if score >= 90 {
		fmt.Println("A등급")
	} else if score >= 80 {
		fmt.Println("B등급")
	} else if score >= 70 {
		fmt.Println("C등급")
	} else if score >= 60 {
		fmt.Println("D등급")
	} else {
		fmt.Println("F등급")
	}
}

// 출력 결과물
// (직접 입력하는 값에 따라 A ~ F 등급이 나옵니다)
