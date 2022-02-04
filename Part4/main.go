package main

import "fmt"

func printSth(msg string) chan string {
	result := make(chan string)
	go func() {
		for i := 0; i <= 5; i++ {
			result <- fmt.Sprintf("%s %d", msg, i)
		}
	}()
	return result
}

func fanIn(chan1, chan2 chan string) chan string {
	res := make(chan string)
	go func() {
		for {
			select {
			case <-chan1:
				res <- <-chan1
			case <-chan2:
				res <- <-chan2
			}
		}
	}()
	return res
}

func main() {
	coffee := printSth("coffee")
	bread := printSth("bread")
	serve := fanIn(coffee, bread)
	for i := 0; i <= 5; i++ {
		fmt.Printf("Receive from %s \n", <-serve)
	}
	fmt.Println("main finished")
}
