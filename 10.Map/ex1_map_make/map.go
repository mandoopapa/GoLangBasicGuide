package main

import "fmt"

func main() {
	Greeting := make(map[string]string)
	Greeting["English"] = "Good morning!"
	Greeting["Français"] = "Bonjour!"
	fmt.Println(Greeting)
}

// 출력 결과물
// map[English:Good morning! Français:Bonjour!]
