package main

import "time"

func write(text string) {
	for {
		println(text)
		time.Sleep(time.Second)
	}
}

func main() {
	go write("Hello from goroutine") // if nescessary, try whitout "go"
	write("This is executed in same time wih goroutine")
}