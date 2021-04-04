package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := &sync.WaitGroup{}
	months := []string{
		"jan", "feb", "mar", "apr", "may", "jun", "jul", "aug", "sep", "oct", "nov", "dec"}

	for i := 0; i < 3; i++ {
	}
	for _, month := range months {
		wg.Add(1)
		go startWorker(month, wg) //кидаем дерьмо лопатами параллельно
	}
	wg.Wait()

}

func startWorker(cnt string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println(cnt, " worked")
}
