package main

import "fmt"

func inversao(numero int) int {
	return numero * -1
}

func inversaoponteiro(numero *int) {
	*numero *= -1
}

func main() {
	n1 := 20
	fmt.Println(n1)
	fmt.Println(inversao(n1))
	fmt.Println(n1)

	n2 := 50
	fmt.Println(n2)
	inversaoponteiro(&n2)
	fmt.Println(n2)
}
