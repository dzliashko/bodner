package calc

import (
	"fmt"
	"strconv"
)

// primitive calc

func add(i, j int) int { return i + j }
func sub(i, j int) int { return i - j }
func mul(i, j int) int { return i * j }
func div(i, j int) int { return i / j }

type opFuncType func(int, int) int

var opMap = map[string]opFuncType{
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

var expressions = [][]string{
	{"2", "+", "3"},
	{"2", "-", "3"},
	{"2", "*", "3"},
	{"2", "/", "3"},
	// {"2", "%", "3"},
	// {"two", "+", "three"},
	// {"5"},
}

func Calc(expression []string) int {
	var result int

	if len(expression) != 3 {
		fmt.Println("invalid expression:", expression)
	}
	p1, err := strconv.Atoi(expression[0])
	if err != nil {
		fmt.Println(err)
	}
	op := expression[1]
	opFunc, ok := opMap[op]
	if !ok {
		fmt.Println("unsupported operator:", op)
	}
	p2, err := strconv.Atoi(expression[2])
	if err != nil {
		fmt.Println(err)
	}
	result = opFunc(p1, p2)

	return result
}

func TestCalc() {
	for i := range expressions {
		result := Calc(expressions[i])
		fmt.Println(result)
	}
}
