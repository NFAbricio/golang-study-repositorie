package main

import "fmt"

func main() {
	usuario1 := map[string]int{ //o que esta nesse [] é a chave, tipo json?
		"Corinthians": 10,
		"Santos":      0,
	}
	fmt.Println(usuario1["Corinthians"])

	usuario2 := map[string]map[string]int{
		"Pessoa": {
			"idade":  1,
			"filhos": 2,
		},
	}
	fmt.Println(usuario2)
	fmt.Println(usuario2["Pessoa"])
}
