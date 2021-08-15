package main

import "fmt"

func main() {
	var multiArray = [2][3][4]int{
		{{1, 2, 3},
			{1, 2, 3},
			{1, 2, 3}},

		{{1, 2, 3},
			{1, 2, 3},
			{1, 2, 3}},
	}
	fmt.Println(multiArray)
}

// 출력결과물
// [[[1 2 3 4] [1 2 3 4] [1 2 3 4]] [[1 2 3 4] [1 2 3 4] [1 2 3 4]]]
