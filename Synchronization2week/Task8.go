package main

import (
	"fmt"
	"sync"
)

// Shutdown Создаем структуру с полем:
type Shutdown struct {
	done chan struct{} // Здесь done - это канал с типом struct{}, главное его преимущество в том что он ничего не весит.
	once sync.Once
}

// NewShutDown Создаем фунцию-конструктор, для реализации нашего обьекта в main().
func NewShutDown() *Shutdown {
	return &Shutdown{done: make(chan struct{})}
}

// Signal создаем метод, как я понял при использовании этого обьекта мы как бы будем отправлять наш сигнал и закрывать канал.
func (s *Shutdown) Signal() {
	s.once.Do(func() {
		close(s.done)
		fmt.Println("Канал Закрыт") // Выполняем какую либо работу.
	})
}

// Check Создаем метод, который будет смотреть закрыт ли канал.
func (s *Shutdown) Check() bool {
	select {
	case <-s.done:
		return true
	default:
		return false
	}
}

func main() {
	NewDone := NewShutDown() // Реализуем обьект через функцию - конструктор
	wg := sync.WaitGroup{}   // Наша любимая синхроннизация

	wg.Add(1)
	go func() {
		defer wg.Done()
		NewDone.Signal() // Создаем горутину в которой посылаем сигнал, итог: канал закрыт
	}()

	// Создаем 10 читателей
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() { // Каждый читатель создает горутину в которой пытается:
			defer wg.Done()
			if NewDone.Check() { // Проверить закрыт ли канал
				fmt.Println("Канал недоступен") // После проверки выводит строку
			}
		}()
	}
	wg.Wait()
}
