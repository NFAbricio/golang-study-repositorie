package tests

import "strings"

//TypeOfAddress checks if the address starts with a valid type and returns it, otherwise returns "Invalid Type"
func TypeOfAddress(adress string) string {
	validTypes := []string{"Rua", "Avenida", "Travessa", "Alameda"}

	firstWord := strings.Split(adress, " ")[0]

	correctType := false
	for _, tipe := range validTypes{
		if tipe == firstWord {
			correctType = true
		}
	}

	if correctType {
		return firstWord
	}

	return "Invalid Type"
}