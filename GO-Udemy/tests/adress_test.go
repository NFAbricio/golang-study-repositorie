package tests

import "testing"

type adressCenary struct {
	adressInsert string
	expectedType string
}

func TestTypeOfAddress(t *testing.T) {
	adressCenary := []adressCenary{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Travessa XYZ", "Travessa"},
		{"Alameda dos Anjos", "Alameda"},
		//{"Rodovia dos Imigrantes", "Invalid Type"},
		//{"Estrada Velha", "Invalid Type"},
	}

	for _, cenary := range adressCenary {
		receivedReturn := TypeOfAddress(cenary.adressInsert)
		if receivedReturn != cenary.expectedType {
			t.Errorf("For the adress %s, expected %s but received %s", cenary.adressInsert, cenary.expectedType, receivedReturn)
		}
	}

}

func TestAnything(t* testing.T) {
	if 1 > 2 {
		t.Error("1 is not greater than 2")
	} // this test is for go test -v
}