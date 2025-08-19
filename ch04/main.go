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

func main() {
	fmt.Println(task01())
}
