package main

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type User struct {
	Name  string `validate:"min=2,max=3"`
	Email string `validate:"required,email"`
}

func validate(val interface{}) error {
	v := reflect.ValueOf(val)
	for i:= 0; i < v.NumField(); i++ {
		field := v.Field(i)
		tag := v.Type().Field(i).Tag.Get("validate")	

		if tag == ""{
			continue
		}

		rules := strings.Split(tag, ",")
		for _, rules := range rules {
			switch {
			case strings.HasPrefix(rules, "min="):
				min, _ := strconv.Atoi(strings.TrimPrefix(rules, "min="))
				if len(field.String()) < min {
					return fmt.Errorf("the field %s must have at least %d characters", v.Type().Field(i).Name, min)
				}
			case rules == "required": 
				if field.String() == "" {
					return fmt.Errorf("the field %s is required", v.Type().Field(i).Name)
				}
			case rules == "email":
				emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
				if !emailRegex.MatchString(field.String()) {
					return fmt.Errorf("the field %s must be a valid email", v.Type().Field(i).Name)
				}
			}
		}
	}
	return nil
}


func main() {
	user := User{
		Name:  "Fabricio",
		Email: "fabricio.com",
	}

	fmt.Println(validate(user))



	t := reflect.TypeOf(user)
	fmt.Println(t)
	fmt.Println("Name->", t.Name())
	fmt.Println("Kind->", t.Kind(),"\n")

	// for i := 0; i < t.NumField(); i++ {
	// 	field := t.Field(i)
	// 	tag := field.Tag.Get("validate")

	// 	//fmt.Println(field)
	// 	//fmt.Println(tag)
	// 	fmt.Println()

	// 	fmt.Printf("%d. %v (%v) %v", i+1, field.Name, field.Type.Name(), tag)
	// }
}