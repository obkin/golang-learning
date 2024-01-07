package main

import "fmt"

type Human struct {
	Name     string
	age      int
	email    string
	password string
}

func main() {
	person := Human{
		Name:     "Yarok",
		age:      21,
		email:    "yarok@gmail.com",
		password: "qwerty1234567890",
	}

	fmt.Println(person)
}
