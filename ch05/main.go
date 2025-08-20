package main

import "fmt"

func prefixer(str string) func(string) string {
	return func(prefix string) string {
		return str + " " + prefix
	}
}

func main() {
	helloPrefix := prefixer("Hello")
	fmt.Println(helloPrefix("Bob"))
	fmt.Println(helloPrefix("Maria"))
}
