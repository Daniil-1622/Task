package main

import "fmt"

// Создаем функцию IntSliceToInteface и передаем туда слайс numbers,а на выходе ждем интерфейс
func IntSliceToInterface(numbers []int) interface{} {
	//Создаем пустой интерфейс для результата
	result := make([]interface{}, len(numbers))
	for i, v := range numbers {
		// Здесь мы записываем ключи в result
		result[i] = v
	}
	//На выходе интерфейс
	return result
}

func main() {
	numbers := []int{1, 2, 3}
	result := IntSliceToInterface(numbers)
	fmt.Println("Результат:", result)
}

//В Golang мы не можем напрямую преобразовать int в interface{}
//из-за того что это разные типы и это привело бы к ошибке копиляции
