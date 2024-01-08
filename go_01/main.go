package main

import "fmt"

type Human struct {
	Name     string
	age      int
	Email    string
	Password string
}

// entry point
func main() {
	// aboutStr()
	// aboutStructures()
	// aboutArr()
	aboutBooleanAndIfElse()
}

// practise

func aboutStr() {
	hello := "Hello"
	world := "World!"

	helloWorld := hello + " " + world
	fmt.Println(helloWorld)
}

func aboutStructures() {
	person := Human{
		Name:     "Yarok",
		age:      21,
		Email:    "kapellwork@gmail.com",
		Password: "qwerty123",
	}

	fmt.Println(person)
}

func aboutArr() {
	var arr [5]int
	arr[1] = 2

	fmt.Println(arr)
}

func aboutBooleanAndIfElse() {
	var alive bool

	alive = false
	alive = true

	if alive {
		fmt.Println("I'm alive!")
	} else {
		fmt.Println("I'm dead.")
	}
}
