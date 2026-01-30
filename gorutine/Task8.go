package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}

	wg.Add(5)
	go func() {
		for i := 1; i < 6; i++ {
			defer wg.Done()
			fmt.Println(i)
			time.Sleep(200 * time.Millisecond)
		}
	}()
	wg.Wait()
}
