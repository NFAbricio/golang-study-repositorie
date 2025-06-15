package main

import "fmt"

func main() {

	func() {
		fmt.Println("Ola")
	}()

	retorno := func(texto string) string {
		return fmt.Sprintf(texto)
	}("eae")

	fmt.Println(retorno)

}
