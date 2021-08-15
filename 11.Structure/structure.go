package main

import "fmt"

type hero struct {
	name   string
	age    int
	weapon string
}

var h hero

func main() {
	h := hero{}
	h.name = "유비"
	h.age = 30
	h.weapon = "자웅일대검"
	fmt.Println(h)
	hero01()
	hero02()
}

func hero01() {
	h.name = "관우"
	h.age = 31
	h.weapon = "청룡언월도"
	fmt.Println(h)
}

func hero02() {
	h.name = "장비"
	h.age = 28
	h.weapon = "장팔사모"
	fmt.Println(h)
}

// 출력 결과물
// {유비 30 자웅일대검}
// {관우 31 청룡언월도}
// {장비 28 장팔사모}
