package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//Создаю константу с максимальным колличеством значений 3
	const MaxWork = 3
	// Создаю буферизованный канал с размером 3 и типом struct{}
	// Как я понял это будет наш сигнал разрешения
	semaphore := make(chan struct{}, 3)
	// Используем синхроннизацию WaitGroup{}
	wg := sync.WaitGroup{}

	// Запускаем 10 иттераций
	for i := 0; i < 10; i++ {
		// На каждый цикл добавляем одну задачу
		wg.Add(1)
		go func(id int) {
			// С помощью defer откладываем выполнения
			defer wg.Done()
			// Получаем сигнал разрешения
			semaphore <- struct{}{}
			// Выводим занятие слота и его id
			fmt.Println("Занял слот", id)
			// Спим 100 мс
			time.Sleep(100 * time.Millisecond)
			// Показываем то что слот свободен и забираем из канала значения
			fmt.Println("Освободил слот")
			<-semaphore
		}(i)
	}
	wg.Wait()
	fmt.Println("Конец")
}
