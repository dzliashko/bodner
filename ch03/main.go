package main

import (
	"fmt"
)

func task01() {
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привет"}
	fmt.Println(greetings)
	fmt.Println(greetings[:2])
	fmt.Println(greetings[1:4])
	fmt.Println(greetings[3:])
}

func task02() {
	var message string = "Hi 👩 and 👨"
	runes := []rune(message)
	fmt.Println(string(runes[3]))
}

func task03() {
	type Employee struct {
		firstName string
		lastName  string
		id        int
	}
	employee1 := Employee{
		"Name1",
		"LastName1",
		1,
	}
	employee2 := Employee{
		firstName: "Name2",
		lastName:  "LastName2",
		id:        2,
	}
	var employee3 Employee
	employee3.firstName = "Name3"
	employee3.lastName = "LastName3"
	employee3.id = 3

	fmt.Println(employee1)
	fmt.Println(employee2)
	fmt.Println(employee3)
}

func main() {
	task03()
}
