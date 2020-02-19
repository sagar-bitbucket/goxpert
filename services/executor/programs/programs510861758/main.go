package main

import "fmt"

func adder(x, y int) int {
	return x + y
}

func main() {
	a := adder(2, 4)
	fmt.Println(a)
}
