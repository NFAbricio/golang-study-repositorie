package main

import "sync"

func write(text string) {
	for i := 0; i < 5; i++ {
		println(text)
	}
}

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(2) //number of goroutines to wait

	go func() {
		write("Hello from goroutine")

		waitGroup.Done() // waitgroup -1
	}()

	go func() {
		write("Another goroutine, but the program dont end")

		waitGroup.Done()
	}()
	
	waitGroup.Wait()
	//wait until all goroutines finish
}