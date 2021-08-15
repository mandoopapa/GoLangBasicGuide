package main

import "fmt"

func main() {
	a := 3
	fmt.Print("a 값 : ")
	switch a {
	case 1:
		fmt.Println("1 입니다")
	case 2:
		fmt.Println("2 입니다")
	case 3:
		fmt.Println("3 입니다")
	}
}

// 출력결과물
// a 값 : 3 입니다
