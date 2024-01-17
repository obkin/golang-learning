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
	// aboutBooleanAndIfElse()
	// aboutDigits()
	aboutSlices()
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
	fmt.Println(arr) // [0,0,0,0,0]

	arr[1] = 2
	fmt.Println(arr) // [0,2,0,0,0]

	arr[1] = 10
	fmt.Println(arr) // [0,10,0,0,0]

	arrWithValues := [3]int{34, 12, 4}
	fmt.Println(arrWithValues) // [34, 12, 4]

	fmt.Println(arrWithValues[2]) // 4
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

func aboutDigits() {
	var num1 int64 = 1234432
	var num2 float32 = 1.24

	var num3 uint16 = 451
	var num4 int16 = 12343

	fmt.Println(num1, num2, num3, num4)
	fmt.Println(int16(num3) + num4)
}

func aboutSlices() {
	week := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(week)      // [0, 1, 2, 3, 4, 5, 6, 7]
	fmt.Println(week[2:6]) // [2, 3, 4, 5]

	weekend := week[6:8]
	fmt.Println(weekend)      // [6, 7] (days, end of the week)
	fmt.Println(week[6:])     // [6, 7] (days, end of the week)
	fmt.Println(len(weekend)) // 2

	some := [...]int{13, 11, 24, 35, 47, 53, 64, 71}
	fmt.Println(some[0:5]) // [13, 11, 24, 35, 47]
}
