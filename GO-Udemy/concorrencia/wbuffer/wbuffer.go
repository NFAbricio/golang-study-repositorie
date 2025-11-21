package main

import "fmt"

func main() {
	channel := make(chan string, 2)
	channel <- "text to send"
	channel <- "another text to uitlize the buffer"
	//channel <- "more text" //this will cause a DEADLOCK, because the buffer is full

	message := <-channel
	message2 := <- channel
	fmt.Println(message)
	fmt.Println(message2)
}