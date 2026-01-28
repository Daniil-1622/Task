package main

import (
	"fmt"
	"time"
)

// Вариант решения 1
func main() {
	// Создаем пустой канал
	ch := make(chan string)
	fmt.Println("Main started")

	go func() {
		// Кладем Значения в канал
		ch <- "Goroutine finished"
		time.Sleep(100 * time.Millisecond)
	}()

	// Выводим значения из канала и ложим в переменную msg
	msg := <-ch
	fmt.Println(msg)
	fmt.Println("Main finished")
}

//Вариант решения 2
//package main
//
//import (
//"fmt"
//"sync"
//"time"
//)
//
//func main() {
// Создаем WaitGroup{}
//	wg := sync.WaitGroup{}
//	fmt.Println("Main started")
// Записываем сколько горутин будет всего
//	wg.Add(1)
//	go func() {
// c помощью wg.Done() будем смотреть когда горутина отработает
//		defer wg.Done()
//		fmt.Println("Goroutine finished")
//		time.Sleep(100 * time.Millisecond)
//	}()
// Расположения горутины в main()
//	wg.Wait()
//	fmt.Println("Main finished")
//}
