package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Davi"
	u.idade = 21
	fmt.Println(u)

	u2 := usuario{"DAVI", 21}
	fmt.Println(u2)

	u3 := usuario{nome: "Richarlison"}
	fmt.Println(u3)
}
