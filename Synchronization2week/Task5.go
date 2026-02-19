/*
Задание:
Банковский счет с RW Mutex
Создайте структуру банковского счета, где:
Проверка баланса происходит очень часто
Изменение баланса (пополнение/списание) происходит редко
Оптимизируйте с помощью sync.RWMutex.
*/

package main

import (
	"fmt"
	"sync"
)

// Создаем структуру BankAccount c полями:
type BankAccount struct {
	balance int          // Это будет наш баланс
	mu      sync.RWMutex // Наш любимый мютекс для блокировки/разблокировки дуступа к данным
}

// Создаем метод Check(), который будет возвращать текущий баланс
func (b *BankAccount) Check() int {
	b.mu.RLock()         //Блокируем доступ чтения  эклюзивных данных
	defer b.mu.RUnlock() // Разблокируем доступ чтения эклюзивных данных
	return b.balance
}

// Создаем второй метод, который будет делать депозит
func (b *BankAccount) AddBalance(amount int) int {
	b.mu.Lock()         //Блокируем доступ к эклюзивным данным
	defer b.mu.Unlock() // Разблокируем доступ к эклюзивным данным
	b.balance += amount // В входные данные метода AddBalance мы вписали amount int это и будет сумма на которую мы хотим пополнить наш баланс
	return b.balance    // возвращаем наш баланс
}

func main() {
	User1 := &BankAccount{
		balance: 100,
	}

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		User1.AddBalance(150)
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			fmt.Println(User1.Check())
		}
	}()
	wg.Wait()
}
