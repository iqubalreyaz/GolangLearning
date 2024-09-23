package main

import "fmt"

func main() {
	var name string
	var age int
	fmt.Printf("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Printf("Enter age: ")
	fmt.Scanln(&age)
	fmt.Printf("\nWelcome %s, Your age is %d.", name, age)

}
