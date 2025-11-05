package main

import "fmt"

func funcao1() {
	fmt.Println("1")
}

func funcao2() {
	fmt.Println("2")
}

func alunoAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada!")
	fmt.Println("Verificando situação do aluno")

	media := n1 + n2/2

	if media >= 6 {
		return true
	}
	return false
}

func main() {
	defer funcao1()
	funcao2()

	fmt.Println(alunoAprovado(5, 4.5))
}
