package main

import "fmt"

func main() {
	fmt.Println(sum(2, 2))
}

func sum(a int, b int) int {
	return a + b
}

func diff(a int, b int) int {
	return a - b
}

func times(a int, b int) int {
	return a * b
}
