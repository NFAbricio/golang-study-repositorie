package main

import "fmt"

func somar(n1 int, n2 int) int {
	return n1 + n2
}

func calculo(n1, n2 int) (int, int) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Print(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := f("print da função 1")
	fmt.Print(resultado)

	_, resultadoSubtracao := calculo(10, 5)
	fmt.Println(resultadoSubtracao)
}
