package main

import (
	"fmt"
	"strings"

	"github.com/iqubalreyaz/GolangLearning/structs/database"
	"github.com/iqubalreyaz/GolangLearning/structs/input"
)

// _ "github.com/iqubalreyaz/GolangLearning/structs/database"

func main() {
	var operation string
	fmt.Printf("Enter operation you want to perform [Entry, Query]: ")
	fmt.Scan(&operation)
	operation = strings.ToLower(operation)
	if operation == "entry" {
		// database.CreateDB()
		data := input.GatherDetails()
		fmt.Println(data)
	} else if operation == "query" {
		data := database.QueryUser()
		fmt.Println(data)
	}
}
