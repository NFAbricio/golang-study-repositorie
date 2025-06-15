package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
}

type estudante struct {
	pessoa
	curso string
	sala  uint8
}

func main() {
	//P1 := pessoa{"Pedor", "Augusto", 21}

	es1 := estudante{pessoa{"", "renato", 22}, "ingles", 14}

	fmt.Println(es1.pessoa.sobrenome)
}
