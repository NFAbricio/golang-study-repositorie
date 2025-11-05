package main

import "fmt"

func fibonnaci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonnaci(posicao-2) + fibonnaci(posicao-1)
}

func main() {
	posicao := uint(12)

	for i := uint(1); i <= posicao; i++ {
		fmt.Print(" ", fibonnaci(i))
	}
}
