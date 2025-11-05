package main

import "fmt"

type animal interface {
	sound() string
}

type Dog struct{}

type Cat struct{}

func (Cat) sound() string {
	return "miau miau"
}

func (Dog) sound() string {
	return "au au"
}

func animalNoise(a animal) {
	fmt.Println(a.sound())
}

func main() {
	dog := Dog{}
	cat := Cat{}
	animalNoise(dog)
	animalNoise(cat)
}

func takeAnimal(a animal) {
	switch t := a.(type){
	case *Dog: //need "*" because methods are pointers receive
		t.sound()
	case *Cat:
	}	
}


//A interface is like a "polimofirms"
//A interface is implemented without a "key word" implement