package main

import (
	"fmt"
	"runtime"
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
	fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())

	go printa()
	go printb()
	go printc()
	time.Sleep(time.Second)
	fmt.Println("\nGo Routine End")
}

// 출력 결과물
// 12 (개발환경마다 상이)
// AAAAAAAAAABBBBBBBBBBCCCCCCCCCC
// Go Routine End
