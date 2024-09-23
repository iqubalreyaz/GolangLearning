package main

import "fmt"

type DependencyFunc func(int) int

func FunctionUnderTest(dep DependencyFunc, num int) int {
	return dep(num) * 2
}

func main() {
	result := FunctionUnderTest(func(num int) int {
		return num + 1 // Mocked dependency
	}, 5)
	fmt.Println(result) // Output: 12
}
