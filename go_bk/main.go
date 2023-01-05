package main

import (
	"fmt"
	"strconv"
	"time"
)

const Py = 3.14

// Auto increment
const (
	A = iota
	B
	C
)

func Plus(x, y int) int {
	return x + y
}

func Div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}
func Double(x int) (result int) {
	result = x * 2
	return
}

func main() {

	var now time.Time = time.Now()
	var value1 int = 100
	var value2 string = "value2"
	var (
		value3 int    = 200
		value4 string = "hello"
	)
	value5 := 400
	value6 := [...]string{"hello", "world"}
	var value7 interface{}

	// str → int
	var value8 string = "1000"
	value9, err := strconv.Atoi(value8)
	if err != nil {
		fmt.Println(err)
	}

	var value10 int = Plus(1, 2)
	var value11, value12 int = Div(3, 4)
	var value13 int = Double(5)

	// ログ出力
	fmt.Println(now)
	fmt.Println(value1)
	fmt.Println(value2)
	fmt.Println(value3)
	fmt.Println(value4)
	fmt.Println(value5)
	fmt.Println(value6)
	fmt.Println(value7)
	fmt.Println(value8)
	fmt.Println(value9)
	fmt.Println(A, B, C)
	fmt.Println(value10)
	fmt.Println(value11)
	fmt.Println(value12)
	fmt.Println(value13)
}
