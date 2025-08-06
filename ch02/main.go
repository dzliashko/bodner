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

func ex03() {
	var b byte = 255
	var smallI int32 = 2_147_483_647
	var bigI uint64 = 18_446_744_073_709_551_615
	b++
	smallI++
	bigI++
	fmt.Println(b)
	fmt.Println(smallI)
	fmt.Println(bigI)
}

func main() {
	ex01()
	ex02()
	ex03()
}
