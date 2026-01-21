package main

import "fmt"

// Создаем интерфейс с методом UpdateName
type Updater interface {
	UpdateName(Name string)
}

// Создаем стуртуру User с полем Name
type User struct {
	Name string
}

// Реализуем метод UpdateName для стуктуры User
// Метод с Указательным Приемником -> Может изменять поле
func (p *User) UpdateName(Name string) {
	p.Name = Name
}

func main() {
	// Создаем обьект User
	User := &User{Name: "Иван"}
	// Присваиваем его переменной интерфейса
	var u Updater = User
	//Меняем поле с помощью метода UpdateName
	u.UpdateName("Катя")
	// Выводим новый обьект
	fmt.Println("Обновленный обьект:", User.Name)
}
