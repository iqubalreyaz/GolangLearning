package main

import (
	"errors"
	"fmt"
)

func handleErrors(fn func() error) error {
	if err := fn(); err != nil {
		return fmt.Errorf("error occurred: %v", err)
	}
	return nil
}

func main() {
	err := handleErrors(func() error {
		// Simulating an operation that may return an error
		return errors.New("something went wrong")
	})
	if err != nil {
		fmt.Println("Error handled:", err)
	} else {
		fmt.Println("No error occurred")
	}
}
