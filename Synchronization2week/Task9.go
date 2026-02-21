package main

import (
	"fmt"
	"sync"
	"time"
)

// Создаем структуру TempsTampStore, у которой есть поля:
type TempsTampStore struct {
	mu      sync.RWMutex // mu - это наша любимая блокировка/разблокировка эклюзивных данных
	Maxsize int          // Maxsize - это число сколько максимум будет в нашем слайсе
	temp    []time.Time  // temp - это слайс с типом времени
}

// Создаем функцию - конструктор которая будет реализововать обьект
func NewTampStore(size int) *TempsTampStore {
	return &TempsTampStore{
		temp:    make([]time.Time, 0, size),
		Maxsize: size,
	}
}

// Создаем метод Add() который будет:
func (t *TempsTampStore) Add(g time.Time) {
	t.mu.Lock()                // Блокировать Mutex
	defer t.mu.Unlock()        // Разблокировать Mutex
	t.temp = append(t.temp, g) // Добавляет в наш временный слайс, промежуток нашего времени

	if len(t.temp) > t.Maxsize { // Проверяет если длинна превышает, максимальное значение
		t.temp = t.temp[1:] // Тогда срезаем первый элемент
	}

}

// Создаем метод Getlast(), который будет возвращать новый промежуток времени
func (t *TempsTampStore) Getlast() []time.Time {
	t.mu.RLock()
	defer t.mu.RUnlock()
	result := make([]time.Time, len(t.temp)) // Здесь мы создаем новый слайс времени, с длинной первого
	copy(result, t.temp)                     // Копируем его в result, тем самым избегая все баги
	return result
}

func main() {
	Result := NewTampStore(3) // Реализуем обьект
	wg := sync.WaitGroup{}

	for i := 0; i < 3; i++ { // Запускаем 3 горутины которые будут добавлять
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			Result.Add(time.Now())
			time.Sleep(time.Second * 3)
		}(i)
	}

	for i := 0; i < 10; i++ { // Запускаем 10 горутин чтения
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			events := Result.Getlast()
			time.Sleep(time.Second * 1)
			for _, t := range events {
				fmt.Printf("%v\n", t.Format("15:04:05.000"))
			}
		}(i)
	}
	wg.Wait()
	fmt.Println("Работа окончена")
}
