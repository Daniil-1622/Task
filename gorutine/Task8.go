package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//Создаем WaitGroup
	wg := sync.WaitGroup{}

	// Записываем сколько горутин будет всего, в нашем случае 5
	wg.Add(5)
	go func() {
		// Проходимся от 1 до 5
		for i := 1; i < 6; i++ {
			// Для каждой итерации wg.Done()
			defer wg.Done()
			// Выводим число
			fmt.Println(i)
			// Делаем паузу 200 мс
			time.Sleep(200 * time.Millisecond)
		}
	}()
	// Расположения нашей горутины в main()
	wg.Wait()
}
