//a pattern for work

package main

import "fmt"

func main() {
	task := make(chan int, 45)
	result := make(chan int, 45)

	go worker(task, result)
	go worker(task, result)
	go worker(task, result)
	//this is for the program run faster

	for i:=0; i<45; i++ {
		task <- i
	}
	close(task)

	for i:=0; i<45; i++ {
		result := <- result
		fmt.Println(result)
	}
	
}

func worker(task <- chan int, result chan<-int) { // task receives, result sends
	for number := range task {
		result <- fibonacci(number)
	}
}

func fibonacci(position int) int{
	if position <= 1{
		return position
	}
	return fibonacci(position-2) + fibonacci(position-1)
}