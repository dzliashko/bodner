package main

import (
	"fmt"
)

func main() {

	// anonymous function
	for i := range 5 {
		func(j int) {
			fmt.Println("printing", j, "from inside of an anonymous function")
		}(i)
	}
}
