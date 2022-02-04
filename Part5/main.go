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

func firstPipeline(nums ...int) chan int {
	result := make(chan int)
	go func() {
		for i := 0; i < len(nums); i++ {
			result <- nums[i]
		}
		close(result)
	}()
	return result
}

func secondPipeline(first chan int) chan int {
	result := make(chan int)
	go func() {
		for item := range first {
			result <- item * item
		}
		close(result)
	}()
	return result
}

func main() {
	first := firstPipeline(1, 2, 3, 4, 5, 6)
	second := secondPipeline(first)
	for i := range second {
		fmt.Printf("Received: %d\n", i)
	}
}
