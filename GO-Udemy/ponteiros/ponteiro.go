package main

import "fmt"

func main() {
	var v1 int  //o valor é 0
	var v2 *int //o valor é nill

	v1, v2 = 10, &v1
	fmt.Println(v1, v2)
	fmt.Println(v1, *v2)

	v1 = 200
	fmt.Println(v1, *v2)

}
