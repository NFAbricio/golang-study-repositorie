package main

import "fmt"

func diadasemana(numero int) string {
	switch numero {
	case 1:
		return "seg"
	case 2:
		return "ter"
	case 3:
		return "quar"
	case 4:
		return "qui"
	case 5:
		return "sex"
	case 6:
		return "sab"
	case 7:
		return "dom"
	default:
		return "não é dia"
	}
}

func main() {
	dia := diadasemana(7)
	fmt.Println(dia)
}
