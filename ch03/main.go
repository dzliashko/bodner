package main

import "fmt"

func task01() {
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привет"}
	fmt.Println(greetings)
	fmt.Println(greetings[:2])
	fmt.Println(greetings[1:4])
	fmt.Println(greetings[3:])
}

func main() {
	task01()
}
