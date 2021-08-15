package main

import "fmt"

func main() {
	a := 3
	b := 5
	c := 6

	fmt.Print("a 값 : ")
	switch a {
	case 1:
		fmt.Println("1 입니다")
	case 2:
		fmt.Println("2 입니다")
	case 3:
		fmt.Println("3 입니다")
	}

	fmt.Print("b 값 : ")
	switch b {
	case 1:
		fmt.Println("1 입니다")
	case 2:
		fmt.Println("2 입니다")
	case 3, 4, 5:
		fmt.Println("3 또는 4 또는 5입니다")
	}

	fmt.Print("c 값 : ")
	switch c {
	case 1:
		fmt.Println("1 입니다")
	case 2:
		fmt.Println("2 입니다")
	case 3, 4, 5:
		fmt.Println("3 또는 4 또는 5입니다")
	default:
		fmt.Println("6 이상입니다")
	}
}

// 출력결과물
// a 값 : 3 입니다
// b 값 : 3 또는 4 또는 5입니다
// c 값 : 6 이상입니다
