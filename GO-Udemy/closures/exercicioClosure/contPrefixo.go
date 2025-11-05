package main

import (
	"fmt"
	"strings"
)

func contadorDePrefixo(prefixo string) func(string) bool {

	verificador := prefixo
	contador := 0
	return func(palavra string) bool {
		if strings.HasPrefix(palavra, verificador) {
			contador += 1
			fmt.Printf("Quantidade de True: %d -> ", contador)
			return true
		} else {
			fmt.Printf("Quantidade de True: %d -> ", contador)
			return false
		}
	}
}

func main() {
	verificadorDeAuto := contadorDePrefixo("auto")

	fmt.Println(verificadorDeAuto("automovel"))
	fmt.Println(verificadorDeAuto("automatizar"))
	fmt.Println(verificadorDeAuto("carro"))
	fmt.Println(verificadorDeAuto("autonomia"))
}
