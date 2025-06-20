package main

import (
	"fmt"
	"strings"
)

func contadorDePrefixo(prefixo string) (func(string) bool, func() int) {
	contagem := 0

	verificador := func(palavra string) bool {
		if strings.HasPrefix(palavra, prefixo) {
			contagem++
			return true
		}
		return false
	}

	contador := func() int {
		return contagem
	}

	return verificador, contador
}

func main() {
	checkAuto, takeCount := contadorDePrefixo("bio")

	fmt.Println(checkAuto("biologia"), "->", takeCount())
}
