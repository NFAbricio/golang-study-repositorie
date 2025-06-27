package main

import "fmt"

//Remember: a Slice just * to an array

// func main() {
// 	arr := [5]int{1, 2, 3, 4, 5}
// 	slc := arr[1:4]
// 	fmt.Println(arr)
// 	arr[2] = 15
// 	fmt.Println(slc)
// 	slc[0] = 123
// 	fmt.Println(arr)
// }

var filmesDB = []string{
	"Poderoso chefao",
	"Titanic",
	"Senhor dos aneis",
	"matrix",
	"Forrest Gump",
	"Rei leao",
	"Harry potter",
	"Gladior",
	"Sexto sentido",
	"O Curioso caso de benjamin button",
	"Pulp fiction",
	"Esqueceram de mim",
	"Exteminador do futro",
	"Batman",
	"Narnia",
}

func main() {
	resultsFromApi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//the cap of a slice double de range
	//var filmes []string

	filmes := make([]string, 0, 10)

	for _, id := range resultsFromApi {
		filme := filmesDB[id]
		fmt.Println(len(filmes), cap(filmes))
		filmes = append(filmes, filme)
		fmt.Println(len(filmes), cap(filmes))
	}
	fmt.Println(filmes)
}
