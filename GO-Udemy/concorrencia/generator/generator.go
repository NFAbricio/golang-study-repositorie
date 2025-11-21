//generator patern

package main

import (
	"fmt"
	"time"
)

func write(text string) <- chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <-fmt.Sprintf("Text received: %s", text)
			time.Sleep(time.Second* 1)
		}
	}()

	return channel
}

func main() {
	channel := write("patern generator")

	for i := 0; i<5; i++ {
		fmt.Println(<- channel)
	}
}