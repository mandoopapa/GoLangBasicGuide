package main

import "fmt"

type SamHero interface {
	Battle()
	Diplomacy
}
type Diplomacy interface {
	Allay()
	Betray()
}
type Hero struct {
}

func (p Hero) Battle() {
	fmt.Println("전투를 개시합니다")
}

func (p Hero) Allay() {
	fmt.Println("동맹을 맺습니다")
}

func (p Hero) Betray() {
	fmt.Println("동맹을 배신합니다")
}

func main() {
	h := Hero{}
	h.Battle()
	h.Allay()
	h.Betray()
}

// 출력 결과물
// 전투를 개시합니다
// 동맹을 맺습니다
// 동맹을 배신합니다
