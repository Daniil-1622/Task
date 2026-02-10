package main

import (
	"sync"
)

// Создаем структуру SafeMap c полями data типа map[string]int и mu типом sync.RWMutex
type SafeMap struct {
	data map[string]int
	mu   sync.RWMutex
}

// Делаем первый метод Set() для стуктуры *SafeMap с рессивером указателем
func (s *SafeMap) Set(key string, val int) {
	// Используем нашу блокировку эклюзивного доступа
	s.mu.Lock()
	// И разблокировку
	defer s.mu.Unlock()
	// Записываем в наш ключ значения
	s.data[key] = val
}

// Создаем второй метод Get для стуктуры *SafeMap с рессивером указателем
func (s *SafeMap) Get(key string) (int, bool) {
	// Используем нашу блокировку чтения эклюзивного доступа
	s.mu.RLock()
	// и разблокировку
	defer s.mu.RUnlock()
	// Тут мы возвращаем значения и проверям есть ли ключ?
	val, exists := s.data[key]
	return val, exists
}
func main() {
	// Создаем Обьект User
	User := &SafeMap{
		// Создаем пустую map
		data: make(map[string]int),
	}
	wg := sync.WaitGroup{}

	// Запускаем цикл на две операции
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			// Используем метод Set() для нашего обьекта
			// id*100 умножает i на 100
			User.Set("key", id*100)
		}(i)
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// Проверяем есть ли ключ?
			if val, ok := User.Get("key"); ok {
				_ = val
			}
		}()
	}
	wg.Wait()
}
