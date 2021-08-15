package main

import "fmt"

func main() {
	a := 1
	b := 2
	fmt.Printf("a = %d, b = %d \n", a, b)
	fmt.Println("호출 후")
	swap(&a, &b)
}

func swap(i *int, j *int) {
	temp := *i
	*i = *j
	*j = temp
	fmt.Printf("a = %d, b = %d \n", *i, *j)
}

// 출력 결과물
// a = 1, b = 2
// 호출 후
// a = 2, b = 1
