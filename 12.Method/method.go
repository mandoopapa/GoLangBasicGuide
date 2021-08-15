package main

import "fmt"

type Account struct {
	balance int
}

func withdrawFunc(a *Account, amount int) {
	a.balance -= amount
}
func (a *Account) withdrawMethod(amount int) {
	a.balance -= amount
}
func main() {
	a := Account{100}
	b := Account{200}
	withdrawFunc(&a, 30)
	withdrawFunc(&b, 30)
	a.withdrawMethod(40)
	b.withdrawMethod(50)
	fmt.Printf("%d, %d", a.balance, b.balance)
}

// 출력 결과물
// 30, 120
