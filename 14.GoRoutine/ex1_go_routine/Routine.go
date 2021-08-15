package main

import (
	"fmt"
	"time"
)

func printa() {
	for i := 0; i < 10; i++ {
		fmt.Print("A")
	}
}

func printb() {
	for i := 0; i < 10; i++ {
		fmt.Print("B")
	}
}

func printc() {
	for i := 0; i < 10; i++ {
		fmt.Print("C")
	}
}

func main() {
	go printa()
	go printb()
	go printc()
	time.Sleep(time.Second)
	fmt.Println("\nGo Routine End")
}

// 출력 결과물
// AAAACAAAAAABBBBBBBBBBCCCCCCCCC (출력시마다 변경됨)
// Go Routine End
