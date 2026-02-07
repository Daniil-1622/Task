package main

import (
	"fmt"
	"sync"
)

// Создаем стуктуру SafeCounter в которую встраивая sync.Mutex для блокировки/разблокировки эклюзивных данных
// А также sync.RWMutex для чтения эклюзивных данных
// Counter - счетчик
type SafeCounter struct {
	// Оставляем только RWMutex
	mr      sync.RWMutex
	Counter int
}

// Создаем метод Inc()
func (s *SafeCounter) Inc() {
	// Блокирует доступ к эклюзивных данных
	s.mr.Lock()
	// Откладывает разблокировку
	defer s.mr.Unlock()
	// Увеличивает счетчик
	s.Counter++
}

// Создаем метод Value()
func (s *SafeCounter) Value() int {
	// Блокируем доступ к чтение эклюзивных данных
	s.mr.RLock()
	// Откладываем разблокировку  чтение эклюзивных данных
	defer s.mr.RUnlock()
	// Возвращаем наш счетчик
	return s.Counter
}
func main() {
	// Реализуем структуру
	counter := SafeCounter{}
	// Создаем WaitGroup
	wg := sync.WaitGroup{}
	// Добавляем 50 задач
	wg.Add(50)

	// Проходимся 50 раз
	for i := 0; i < 50; i++ {
		// Запускаем gorutine
		go func() {
			// На каждой интерации отнимаем от wg.Add(50) еденицу
			defer wg.Done()
			// Используем наш метод Inc() к обькту counter
			counter.Inc()
		}()
	}
	// Ждем выполнения всех задач
	wg.Wait()
	// Выводим
	fmt.Println("Результат:", counter.Value())
}

// Это абсолютно безопасно
// В нашем случае метод Inc() может только писать, а метод Value() только читать
// Запись lock() исключает одновременного чтения и записи, а чтение RLock позволяет множеству читателей работать парралельно
// Полностью исключает DataRace
