package main

import "fmt"

type user struct {
	name string
	age  int
}

func (u user) presentation() {
	fmt.Printf("Hello, my name is %s\n", u.name)
}

func maiority(u user) bool {
	return u.age >= 18
}

func (u *user) birthday() {
	u.age++
}

func main() {
	user1 := user{"antony", 18}
	fmt.Println(maiority(user1))
	user1.presentation()
	user1.birthday()
	fmt.Println(user1)

}
