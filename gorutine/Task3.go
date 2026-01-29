// Вариант решения 1
package main

import (
	"fmt"
	"sync"
	"time"
)

// Создаем функцию PrintMsg в которую передаем id int и Указатель на WaitGroup
func PrintMsg(id int, wg *sync.WaitGroup) {
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
		go PrintMsg(i, wg)
		// Даем время на выполнение
		time.Sleep(100 * time.Millisecond)
	}
	wg.Wait()
}

// Вариант решения 2

//package main
//
//import (
//"fmt"
//"time"
//)
//// Создаем функцию PrintMsg в которую передаем i
//func PrintMsg(i int) {
//	// Выводим с помощью Println()
//	fmt.Println("Gorutine: ", i)
//}
//func main() {
//	// Запускаем цикл от 1 до 4
//	for i := 1; i < 4; i++ {
//		// Запускаем gorutine на функцию PrintMsg
//		go PrintMsg(i)
//	}
//	// Даем время на выполнение
//	time.Sleep(100 * time.Millisecond)
//}
