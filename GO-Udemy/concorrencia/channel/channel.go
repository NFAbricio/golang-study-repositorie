package main

import (
	"fmt"
	"time"
)

func write(text string, channel chan string) {
	for i := 0; i < 5; i++ {
		channel <- text
		time.Sleep(time.Second)
	}

	close(channel) //whiout this, the for in main will never end, so this cause a DEADLOCK
}

func main() {
	channel := make(chan string) //make is for creating and initializing channels, maps and slices

	go write("Hello from goroutine", channel)

	//message2 := <-channel //receive message from channel, dont need this line if using for range, just for example how pass a channel to a variable
	
	for message := range channel {
		fmt.Println(message)
	}
}