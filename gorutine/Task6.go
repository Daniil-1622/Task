package main

import "fmt"

func main() {
	// Вариант решения номер один

	// Создаем канал для нашего сигнала
	done := make(chan struct{})

	go func() {
		// Выводим "Worker done"
		fmt.Println("Worker done")
		// Закрываем наш канал
		close(done)
	}()

	// Получаем наш сигнал
	fmt.Println(<-done)

	// Можем дальше двигаться по main() и вывести "Main continues"
	fmt.Println("Main continues")
}

// Вариант решения номер два
//package main
//
//import "fmt"
//
// Создаем функцию Worker в которую передаем канал со стуктурой, это и будет наш сигнал
//func Worker(done chan struct{}) {
// Выводим "Worker done"
//fmt.Println("Worker done")
//
// Отправка нашего сигнала
//done <- struct{}{}
//}
//
//func main() {
//	// Создаем пустой канал для нашего сигнала
//	done := make(chan struct{})
//
//	// Запускаем нашу горутину на функцию Worker
//	go Worker(done)
//	// Принятие нашего сигнала, без его маин дальше не пойдет
//	<-done
//
//	// Вывод "Main continues"
//	fmt.Println("Main continues")
//}
