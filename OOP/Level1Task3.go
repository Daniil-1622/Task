package main

import "fmt"

type bankAccount struct {
	balance float64
}

func (b *bankAccount) Deposit(amount float64) {
	if amount >= 0 {
		b.balance = amount + b.balance
	}
}

func (b bankAccount) GetBalance() float64 {
	return b.balance
}

func main() {
	User := bankAccount{300}
	fmt.Println("До депозита:", User.GetBalance())

	User.Deposit(200)
	fmt.Println("После депозита:", User.GetBalance())
}

//Почему нельзя напрямую изменить `balance` извне пакета?
//Потому что поле 'balance' начинается с маленькой буквы и является приватным полем, значит что мы бы его не увидели в другом пакете
