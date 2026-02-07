// Вариант решения 1
package main

import (
	"fmt"
	"sync"
)

// Создаем стуктуру SafeCounter, со внутренним счетчиком Counter
type SafeCounter struct {
	Counter int
}

// Создаем Метод Inc использует указательный приёмник
func (s *SafeCounter) Inc() {
	// Увеличиваем наш счетчик на 1
	s.Counter++
}

// Создаем метод Value() возвращает текущее значение счётчика
func (s SafeCounter) Value() int {
	return s.Counter
}
func main() {
	// Создаем Mutex для блокировки/разблокировки экслюзивных данных
	mu := sync.Mutex{}
	// Создаем RWMutex для блокировки/разблокировки чтения эклюзивных данных
	mr := sync.RWMutex{}
	// Создаем WaitGroup для синхронизации
	wg := sync.WaitGroup{}
	// Добавляем 2 задачи нашей WaitGroup
	wg.Add(2)

	// Реализуем нашу структуру SafeCounter
	Admin := SafeCounter{0}

	// Запускаем нашу gorutine для увеличения нашего счетчика
	go func() {
		// Откладываем wg.Done() в нашей gorutine, после того как он отработает от wg.Add(2) отнимут единицу
		defer wg.Done()
		// Проходимся от 0 до 50
		for i := 0; i < 50; i++ {
			// Блокируем доступ к нашим эклюзивным данным
			mu.Lock()
			// Применям метод Inc() к нашему обькту Admin
			Admin.Inc()
			// Разблокируем доступ к эклюзивным данным
			mu.Unlock()
		}
	}()
	// Запускаем вторую gorutine для чтения нашего счетчика
	go func() {
		// Откладываем wg.Done() в нашей gorutine, после того как он отработает от wg.Add(1) отнимут единицу
		defer wg.Done()
		// Проходимся от 0 до 50
		for i := 0; i < 50; i++ {
			// Блокируем доступ к чтения наших эклюзивных данных
			mr.RLock()
			// Применяем наш метод Value() к обьекту Admin
			Admin.Value()
			// Разблокируем доступ к чтения наших эклюзивных данных
			mr.RUnlock()
		}
	}()
	// Ждем пока wg.Add() будет 0
	wg.Wait()
	// Выводим наш обьект после добавления и чтения
	fmt.Println("Обьект:", Admin)
}

// Вариант решения 2, после первого решения решил посмотреть как можно еще решить эту задачу
//package main
//
//import (
//"fmt"
//"sync")
// Создаем стуктуру SafeCounter в которую встраивая sync.Mutex для блокировки/разблокировки эклюзивных данных
//А также sync.RWMutex для чтения эклюзивных данных
//Counter - счетчикtype SafeCounter struct {
//mu      sync.Mutex
//mr      sync.RWMutex
//Counter int
//}

// Создаем метод Inc()
//func (s *SafeCounter) Inc() {
// Блокирует доступ к эклюзивных данных
//s.mu.Lock()
// Откладывает разблокировку
//defer s.mu.Unlock()
// Увеличивает счетчик
//s.Counter++
//}
// Создаем метод Value()
//func (s SafeCounter) Value() int {
// Блокируем доступ к чтение эклюзивных данных
//s.mr.Lock()
// Откладываем разблокировку  чтение эклюзивных данных
//defer s.mr.Unlock()
// Возвращаем наш счетчик
//return s.Counter
//}
//func main() {
// Реализуем структуру
//counter := SafeCounter{}
// Создаем WaitGroup
//wg := sync.WaitGroup{}
// Добавляем 50 задач
//wg.Add(50)

// Проходимся 50 раз
//for i := 0; i < 50; i++ {
// Запускаем gorutine
//go func() {
// На каждой интерации отнимаем от wg.Add(50) еденицу
//defer wg.Done()
// Используем наш метод Inc() к обькту counter
//counter.Inc()
//}()
//}
// Ждем выполнения всех задач
//wg.Wait()
// Выводим
//fmt.Println("Результат:", counter)
//}
