package main

import "fmt"

type List struct {
	data []int
}

func (l *List) Len() int {
	if l == nil {
		return 0
	} else {
		return len(l.data)
	}
}

func main() {
	var list *List
	fmt.Println(list.Len())

	list = &List{data: []int{1, 2, 3}}
	fmt.Println(list.Len())
}
