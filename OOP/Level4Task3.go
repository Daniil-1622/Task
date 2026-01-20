package main

import (
	"context"
	"fmt"
)

type Worker interface {
	Do(ctx context.Context) error
}

type FileProcessor struct {
	filename string
}

func (f FileProcessor) Do(ctx context.Context) error {
	fmt.Println("Обрабатываю файл:", f.filename)
	return nil
}

type NetworkFetcher struct {
	url string
}

func (n NetworkFetcher) Do(ctx context.Context) error {
	fmt.Println("Запрашиваю URL с:", n.url)
	return nil
}

func RunWorker(w Worker) {
	ctx := context.Background()
	err := w.Do(ctx)
	if err != nil {
		fmt.Println("Ошибка")
	}
}

func main() {
	RunWorker(FileProcessor{"data.txt"})
	RunWorker(NetworkFetcher{"https://img.freepik.com/free-photo/fluffy-kitten-playing-looking-charming-eyes-curious-nature-beauty-generated-by-artificial-intelligence_188544-241336.jpg?semt=ais_hybrid&w=740"})
}
