package main

import (
	"fmt"
	"strings"
)

func main() {
	Len()
	Cap()
}

// // 	// sl1 := make([]string, 10)
// // 	// sl2 := make([]string, 10, 20)

// // 	// fmt.Println("sl1 - ", len(sl1), cap(sl1))
// // 	// fmt.Println("sl2 - ", len(sl2), cap(sl2))

// // 	// sl2 = sl2[2:]
// // 	// fmt.Println("sl2 - ", len(sl2), cap(sl2))

var letters = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w",
	"x", "y", "z"}

func Len() []string {

	sl1 := make([]string, len(letters))
	for key, letter := range letters {
		sl1[key] = letter
	}

	for _, letter := range letters {
		sl1 = append(sl1, strings.ToUpper(letter))
	}

	fmt.Println("sl1 - ", len(sl1), cap(sl1))

	return sl1
}

func Cap() []string {
	sl2 := make([]string, len(letters), len(letters)*2)
	for key, letter := range letters {
		sl2[key] = letter
	}

	for _, letter := range letters {
		sl2 = append(sl2, strings.ToUpper(letter))
	}

	fmt.Println("sl2 - ", len(sl2), cap(sl2))

	return sl2

}
