package main

import (
	"ch05/closures"
	"fmt"
)

func main() {
	twoBase := closures.MakeMult(2)
	threeBase := closures.MakeMult(3)
	for i := 0; i < 3; i++ {
		fmt.Println(twoBase(i), threeBase(i))
	}
}
