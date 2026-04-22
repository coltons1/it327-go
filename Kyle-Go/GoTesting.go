package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(addToNested(5))
}

func addToNested(a int) int {
	addVal := func(x, y int) int {
		return x + y
	}
	return addVal(5, 10) + a
}
