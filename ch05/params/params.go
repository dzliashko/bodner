package params

import "fmt"

type person struct {
	age  int
	name string
}

func modifyFails(i int, s string, p person) {
	i *= 2
	s = "Goodbye"
	p.name = "Bob"
}

func Params() {
	p := person{}
	i := 2
	s := "Hello"
	modifyFails(i, s, p)
	fmt.Println(i, s, p)
}

func ModifyMap(m map[int]string) {
	m[2] = "hello"
	m[3] = "goodbye"
	delete(m, 1)
}

func ModifySlice(s []int) {
	for k, v := range s {
		s[k] = v * 2
	}
	s = append(s, 10)
}
