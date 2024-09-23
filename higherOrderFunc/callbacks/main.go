package main

import "fmt"

func Process(list []int, callback func(int)) {
	for _, item := range list {
		callback(item)
	}
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	Process(numbers, func(num int) {
		fmt.Println(num * 2)
	})
}
