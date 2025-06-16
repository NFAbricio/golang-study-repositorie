package main

import "fmt"

func recuperandoDefer() {
	fmt.Println("alguma tentativa")
}

func alunoAporvado(n1, n2 float64) bool {
	defer recuperandoDefer()

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("função recuparada apesar do panic")
		}
	}()

	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("Média exatamente igual a 6")
}

func main() {
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println("função recuparada apesar do panic")
	// 	}
	// }()
	fmt.Println(alunoAporvado(6, 6))
	fmt.Println("Termino")
}
