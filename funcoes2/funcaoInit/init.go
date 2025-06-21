package main

import "fmt"

var n int

func init() {
	fmt.Println("init function")
	n = 10
}

func main() {
	fmt.Println("main function")
	fmt.Println(n)
}
