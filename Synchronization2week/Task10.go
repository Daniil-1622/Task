package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Создаем структуру RateLimiter, с полями:
type RateLimiter struct {
	call  atomic.Int64 // Наш счетчик
	limit int64        // Лимит вызовов счетчика
}

// Создаем метод CheckCall(), который будет увеличивать счетчик, до того момента пока не будет привышен лимит
func (r *RateLimiter) CheckCall() bool {
	result := r.call.Add(1) // В переменную result, добавляем атомарно увеличенное число на еденицу

	if result <= r.limit { // Проверяем счетчик не больше лимита
		fmt.Println("Счетчик увеличен") // Печатаем текст
		return true                     // Возвращаем true
	}
	// Как счетчик будет больше лимита, отнимаем еденицу, т.к я понял будет 11
	r.call.Add(-1)
	// И возвращаем false
	return false
}

// Создаем второй метод Timer(), который будет фоново обновлять наш счетчик раз в 20 секунд
func (r *RateLimiter) Timer() {
	// В методе запускаем горутину
	go func() {
		ticker := time.NewTicker(20 * time.Second) // Создаем ticker
		defer ticker.Stop()                        // Обязательно освобождаем ресурсы

		for {
			<-ticker.C     // Ждём сигнал (раз в 20 секунд)
			r.call.Swap(0) // Выполняем задачу, которая обнуляет наш счетчик
		}
	}()
}

func main() {
	// Реализуем обьект и ставим лимит - 10
	limiter := &RateLimiter{
		limit: 10,
	}

	wg := sync.WaitGroup{} // Любимая синхроннизация

	for i := 0; i < 11; i++ { // Создаем 11 горутин
		wg.Add(1)
		go func() {
			defer wg.Done()
			limiter.CheckCall() // В каждой вызываем метод CheckCall()
		}()
	}

	limiter.Timer() // Запускаем наш фоновый таймер
	wg.Wait()

	fmt.Println(" Ждем 25 секунд для проверки сброса")
	time.Sleep(25 * time.Second)

	fmt.Println("\nПроверка после сброса:")
	limiter.CheckCall()
	limiter.CheckCall()
}
