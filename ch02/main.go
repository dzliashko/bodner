package main

import "fmt"

func ex01() {
	i := 20
	var f float64 = float64(i)
	fmt.Println(i)
	fmt.Println(f)
}

func ex02() {
	const value = 15
	i := value
	f := value
	fmt.Println(i)
	fmt.Println(f)
}

func main() {
	ex01()
	ex02()
}
