package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	// Создаем переменную flag с типом int32, это наш будущий флажок
	var flag int32
	// Используем WaitGroup для синхроннизации горутин
	wg := sync.WaitGroup{}
	wg.Add(2)

	// В первой горутине мы будем запускать бесконечный цикл, который будет проверять значение флага
	go func() {
		defer wg.Done()
		for {
			if atomic.LoadInt32(&flag) == 1 { // Если значение flag изменится на еденицу, тогда бесконечный цикл остановится
				break
			}
			if atomic.LoadInt32(&flag) == 0 { // А если никто не изменит значение flag, тогда программа будет спать одну секунду
				time.Sleep(time.Second) // и выводить слово "Работаем"
				fmt.Println("Работаем")
			}
		}
	}()

	// Во второй горутине мы будем менять состояние flag
	go func() {
		defer wg.Done()                         // Тоесть наша горутина спит 5 секунд
		time.Sleep(5 * time.Second)             // А после с помощью функции CompareAndSwapInt32 меняет значение flag на еденицу
		atomic.CompareAndSwapInt32(&flag, 0, 1) // И наш бесконечный цикл остановится и перестанет выводить слово "Работаем"
	}()
	wg.Wait()
	fmt.Println("Флаг изменен")
}
