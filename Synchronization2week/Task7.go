package main

import (
	"fmt"
	"sync"
	"time"
)

// Создаем кэш статистических данных
type Caches struct {
	data map[string]int
	mu   sync.RWMutex
}

// Реализуем метод Set который будет записывать данные в нашу мапу
func (c *Caches) Set(key string, val int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = val // Ключевое действие
}

// Реализуем метод Get который будет читать наш ключ и выводить значение
func (c *Caches) Get(key string) int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.data[key]
}

// Самый лучший метод Update
func (c *Caches) Update(NewMap map[string]int) {
	temp := make(map[string]int) // Создаем временную мапу
	for k, v := range NewMap {   // Проходимся по входной и ложим в ключи временной мапы значения
		temp[k] = v
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	c.data = temp // Меняем ссылку, т.к мы подготовили временную мапу, datarace не будет
}

func main() {
	// Реализуем обьект
	Service := &Caches{
		data: make(map[string]int),
	}
	wg := sync.WaitGroup{} // Используем WaitGroup для синхронизации

	wg.Add(1) // Для метода Set добавляем одну задачу WaitGroup
	go func() {
		defer wg.Done()
		Service.Set("key", 77) // Вписываем в нашу мапу ключ - значения
	}()

	// Запускаем самую интересную горутину
	go func() {
		// Используем ticker
		ticker := time.NewTicker(time.Minute)
		// Откладываем остановку
		defer ticker.Stop()

		// ticker даст сигнал бескочнему циклу, когда пройдет минута и он выполнит следующие действие
		for {
			<-ticker.C
			newMap := map[string]int{"key": 100, "hits": 500} // Создаст и заполнит новую мапу
			Service.Update(newMap)                            // Обновит старую на новую
		}
	}()

	// Запускаем 10 горутин
	for i := 0; i < 10; i++ {
		// На каждую даем по одной задаче
		wg.Add(1)
		go func() {
			// В самой горутине откладываем выполнения
			defer wg.Done()
			for j := 0; j < 100; j++ { // В каждой горутине будет 10 читателей
				fmt.Println(Service.Get("key")) // Используем метод Get для чтения ключа "key"
				time.Sleep(100 * time.Millisecond)
			}
		}()
	}

	wg.Wait()
	fmt.Println("Ждем обновления")
	time.Sleep(65 * time.Second)
	fmt.Println(Service)
	fmt.Println("Программа завершена")
}
