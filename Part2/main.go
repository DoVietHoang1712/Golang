package main

import (
	"fmt"
	"sync"
)

func printNumber(wg *sync.WaitGroup, chanNumber chan int) {
	result := 0
	for i := 0; i <= 100; i++ {
		result += i
	}
	chanNumber <- result
	wg.Done()
}

func printChar(wg *sync.WaitGroup, chanChar chan string) {
	result := ""
	for i := 'A'; i < 'A'+26; i++ {
		result += string(i)
	}
	chanChar <- result
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	chanNumber := make(chan int)
	chanChar := make(chan string)
	wg.Add(2)
	go printNumber(&wg, chanNumber)
	go printChar(&wg, chanChar)
	fmt.Println("KQNumber", <-chanNumber)
	fmt.Println("KQChar", <-chanChar)
	wg.Wait()
}
