package main

import (
	"fmt"
	"time"
)

func main() {
	// Создаем канал
	done := make(chan bool)

	go func() {
		// Создаем тикер с интервалом 300 мс
		ticker := time.NewTicker(300 * time.Millisecond)
		// С помощью defer, стопаем тикер, что бы не было утечки
		defer ticker.Stop()

		// Заходим в бесконечный цикл
		for {
			// Входим в select один раз
			select {
			// Прошло 300 мс и в канал пришел флажок, что пора работать
			case <-ticker.C:
				// Выводит Tick каждый 300 мс
				fmt.Println("Tick")
			case <-done:
				// После того как мы поставили true, программа видит что done изменился
				// И выходит с помощью return из for
				return
			}
		}
	}()
	// Блокируем основной поток программы благодаря time.Sleep
	// Что бы он прежде временно не сделал true и select успел отработать
	time.Sleep(1500 * time.Millisecond)
	// Меняем в true, что бы select обнаружил изменения в done и сделал return
	done <- true
	// Дальше работаем по main()
	fmt.Println("Stopped")
}
