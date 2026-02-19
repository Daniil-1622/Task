package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

// Создаем структуру Play с полем id - уникальных идентификаторов типа uint64
type Play struct {
	id uint64
}

// Создаем метод NextId, который будет увеличивать число на еденицу возвращать число типа uint64
func (p *Play) NextId() uint64 {
	return atomic.AddUint64(&p.id, 1) // Возвращаем &p.id - наше текущее число, а delta 1 - это число на которые мы хотим увеличить id
}

func main() {
	Player := &Play{ // Реализуем обьект
		id: 0,
	}

	for i := 0; i < 10; i++ { // Запускаем 10 иттераций
		fmt.Println(Player.NextId()) // Применяем наш метод к нашему обьекту 10 раз
	}

	time.Sleep(time.Second) // Даем время горутине
}
