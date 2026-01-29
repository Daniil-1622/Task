package main

import (
	"fmt"
	"sync"
)

// Создаем функцию PrintMsg в которую передаем id int и Указатель на WaitGroup
func PrintMsg1(id int, wg *sync.WaitGroup) {
	// Откладываем завершения WaitGroup
	defer wg.Done()
	// Выводим надпись ("Goroutine: &d", id)
	fmt.Println("Goroutine: ", id)
}
func main() {
	// Разыменовываем WaitGroup, для того что бы работать с оригиналом
	wg := &sync.WaitGroup{}

	// Проходимся циклом for от 1 до 4, с i++
	for i := 1; i < 4; i++ {
		// Каждую интерацию добавляем еденицу к счетчику
		wg.Add(1)
		// Запускаем нашу gorutine
		go PrintMsg1(i, wg)
	}
	wg.Wait()
}
