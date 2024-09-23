package main

import "fmt"

func pointers() {
	Firstnum := 100
	numRef := &Firstnum

	updateNum(numRef)
}

func updateNum(numRef *int) {
	*numRef = 100
	fmt.Println(*numRef)
}
