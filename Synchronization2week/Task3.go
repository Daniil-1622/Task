package main

import (
	"fmt"
	"sync"
)

// Для начала скажу что Конфигурация - это параметры программы, которые определяют поведение программы без изменения исходного кода

// Создаем структуру с именем Config с полями:
type Config struct {
	data map[string]int // data - это мапа для хранения данных
	mu   sync.RWMutex   // Mutex для блокировки/разблокировки доступа к данным
}

// Создаем метод Get()
func (c *Config) Get(key string) int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.data[key] // Он возвращает наш ключ
}

// Дальше создаем метод NewMap он будет делать снимок, как выглядит наша мапа сейчас
func (c *Config) NewMap() map[string]int {
	safe := make(map[string]int)
	c.mu.RLock()
	for k, value := range c.data { // Проходимся по нашей мапе data
		safe[k] = value // Записываем ключ в нашу временную мапу safe
	}
	c.mu.RUnlock()

	return safe // Возвращаем safe
}

// Создаем третий метод, он будет записывать данные в нашу мапу
func (c *Config) Set(key string, value int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value // Тут мы ложим в наш ключ = значение мапы
}

// Метод Reload будет перезагружать нашу мапу
func (c *Config) Reload(NewMap map[string]int) {
	// Тоесть мы создаем временную мапу temp
	temp := make(map[string]int)
	for k, v := range NewMap { // Копируем входную мапу во временную
		temp[k] = v
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	c.data = temp // Меняем ссылку, тоесть с.data теперь temp

}

func main() {
	// Реализуем обьект
	Service := &Config{
		data: make(map[string]int),
	}

	// Используем синхроннизацию
	wg := sync.WaitGroup{}
	wg.Add(4)

	// Запускаем метод Set() - он записывает значения в нашу пустую мапу
	go func() {
		defer wg.Done()
		Service.Set("key", 32)
	}()

	// Запускаем метод чтения Get(), который будет читать нашу мапу 30 раз
	go func() {
		defer wg.Done()
		for i := 0; i < 30; i++ {
			Service.Get("key")
		}
	}()

	// Обновляем мапу
	go func() {
		defer wg.Done()
		NewCfg := map[string]int{"key": 44} // Реализуем новую мапу NewCfg и записываем новую пару ключ - значение
		Service.Reload(NewCfg)
	}()

	//Выводит что находится в нашей мапу в данный момент,  как бы скриншот 3 раза
	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			fmt.Println(Service.NewMap())
		}
	}()
	wg.Wait()
	fmt.Println("Старая:", Service.data)
}
