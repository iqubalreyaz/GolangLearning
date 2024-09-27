package main

import "fmt"

func main() {

	fmt.Println("Result is : ")
}

func divide(deno int) (int, error) {
	defer func() (int, error) {
		return 0, nil
	}()

	fmt.Println("Enter denominator: ")
	fmt.Scan(&deno)
	res := 100 / deno
	return res, nil
}
