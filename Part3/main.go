package main

import "fmt"

func printSth(msg string) chan string {
	result := make(chan string)
	go func() {
		for i := 0; i < 5; i++ {
			result <- fmt.Sprintf("%s %d", msg, i)
		}
	}()
	return result
}

func main() {
	bridge := printSth("hello")
	for i := 0; i < 5; i++ {
		fmt.Printf("Receive from %s \n", <-bridge)
	}
	fmt.Println("main finished")
}
