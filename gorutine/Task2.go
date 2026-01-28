package main

import (
	"fmt"
	"sync"
)

// Создаем фунцию greet, которая выводит Hello + name
func greet(name string) {
	fmt.Println("Hello, ", name)
}

func main() {
	// Создаем WaitGroup{}
	wg := sync.WaitGroup{}
	// Записываем сколько горутин будет использоваться
	wg.Add(1)
	go func() {
		//Done() покажет когда горутина отработает
		defer wg.Done()
		// Используем нашу функцию greet с параметром "Alica"
		greet("Alica")
	}()
	// Показываем где будет вывод
	wg.Wait()
}
