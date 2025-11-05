// desecourage use "i" in interface name, like "iAnimal"
// end qith "er"
package main

import (
	"fmt"
)

type Pessoa struct {}

func (Pessoa) String() string {
	return "Hello, my name is Fabricio"
}

func generic(interf interface{}) {
	fmt.Println(interf)
	//a interface for generic types 
}

func main() {
	p := Pessoa{}
	fmt.Println(p)

	generic(1)
	generic("Corinthians")
	generic(true)
	
}