package main

import "fmt"

func main() {
	array1 := [5]int{10, 20, 30}
	fmt.Println(array1)

	var array2 [2]int
	array2[1] = 2
	fmt.Println(array2)

	//array tamanho fixo

	array3 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(array3)
	//array3[6] = 10
	//os "..." se baseiam no quanto vc colocou dentro do array na primeira vez

	slice := []int{90, 40, 50}
	fmt.Println(slice)

	slice2 := array3[1:4]
	fmt.Println(slice2)
	array3[3] = 111
	fmt.Println(slice2)

	//slice sempre se referencia-se a um array
}
