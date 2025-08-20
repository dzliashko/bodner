package main

import (
	"ch05/params"
	"fmt"
)

func main() {
	params.Params()

	m := map[int]string{
		1: "first",
		2: "second",
	}
	params.ModifyMap(m)
	fmt.Println(m)

	s := []int{1, 2, 3}
	params.ModifySlice(s)
	fmt.Println(s)

}
