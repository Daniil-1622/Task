package main

import (
	"fmt"
	"sync"
)

func main() {
	// Создаем глобальную переменную counter типа int
	var counter int
	// Создаем Mutex для защиты эксклюзивных данных
	mu := sync.Mutex{}
	// Создаем WaitGroup для синхронизации gorutine
	wg := sync.WaitGroup{}

	// Добавляем 100 задач нашей WaitGroup
	wg.Add(100)
	//Запускаем gorutine
	go func() {
		// Проходимся от 0 до 100
		for i := 1; i <= 100; i++ {
			// Отнимаем от wg.Add(100) каждую итерацию по 1
			defer wg.Done()
			// Блокируем доступ к эклюзивным данным
			mu.Lock()
			// Критическая часть, записываем в наш счетчик число i
			counter = i
			// Разблокируем доступ к эклюзивным данным
			mu.Unlock()
		}
	}()
	// Ждем пока wg.Add(100) будет равным 0
	wg.Wait()
	// Выводим наш счетчик
	fmt.Println("Счетчик:", counter)
}
