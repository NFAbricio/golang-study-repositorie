package main

import (
	"fmt"
	"time"
)

func write(text string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprint("Text received: ", text)
			time.Sleep(time.Second * 1)
		}
	}()

	return channel
}

func multiplexator(input1, input2 <- chan string) <- chan string {
	exitChannel := make(chan string)

	go func() {
		for {
			select{
			case message := <- input1:
				exitChannel <- message
			case message := <- input2:
				exitChannel <- message
			}
		}
	}()

	return exitChannel
}

func main() {
	channel := multiplexator(write("channel 1"), write("channel 2"))

	for i := 0; i<10; i++ {
		fmt.Println(<- channel)
	}
}