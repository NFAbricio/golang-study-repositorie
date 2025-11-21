package main

import (
	"fmt"
	"time"
)

func main() {
	channel1, channel2 := make(chan string), make(chan string)

	go func() {
		for{
			time.Sleep(time.Millisecond * 500)
		channel1 <- "channel 1"}
	}()

	go func() {
		for {time.Sleep(2 * time.Second)
		channel2 <- "channel 2"}
	}()

	for {
		select { //select waits for one of the channels to be ready, whitout select, the first channel after is printed, he will wait the second channel to be ready
			case msg1 := <- channel1:
				fmt.Println(msg1)
			case msg2 := <- channel2:
				fmt.Println(msg2)
		}
	}


}