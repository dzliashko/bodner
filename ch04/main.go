package main

import (
	"fmt"
	"math/rand"
)

func task01() []int {
	randomInts := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		randomInts = append(randomInts, rand.Intn(100))
	}
	return randomInts
}

func task02(randomInts []int) {
	for _, v := range randomInts {
		switch {
		case v%2 == 0 && v%3 == 0:
			fmt.Println("Шесть!")
		case v%2 == 0:
			fmt.Println("Два!")

		case v%3 == 0:
			fmt.Println("Три!")
		default:
			fmt.Println("Неважно")
		}

	}
}

func main() {
	randomInts := task01()
	fmt.Println(randomInts)
	task02(randomInts)
	var total int
	for i := 0; i <= 10; i++ {
		total := total + i
		fmt.Println(total)
	}
	fmt.Println(total)
}
