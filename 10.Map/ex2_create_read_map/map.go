package main

import "fmt"

func main() {
	// var iMap map[int]int
	iMap := make(map[int]int)
	iMap[0] = 3
	iMap[5] = 9
	val, suc := iMap[3]
	fmt.Println(iMap[0], iMap[5], iMap[3])
	fmt.Println(val, suc)
}

// 출력 결과물
// 3 9 0
// 0 false
