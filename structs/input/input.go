package input

import (
	"fmt"
	"strings"

	"github.com/iqubalreyaz/GolangLearning/structs/data"
	"github.com/iqubalreyaz/GolangLearning/structs/database"
	"github.com/iqubalreyaz/GolangLearning/structs/validation"
)

var (
	Name     string
	Mobile   int
	Email    string
	Locality string
	State    string
	Pin      int
	Country  string
)

func GatherDetails() (detail data.ID) {
	// Map countries to their valid states
	validStates := map[string][]string{
		"INDIA":  {"TELANGANA", "MAHARASHTRA", "KARNATAKA"},
		"USA":    {"TEXAS"},
		"UK":     {"LONDON"},
		"CANADA": {"TORONTO"},
	}
	fmt.Print("Enter name: ")
	fmt.Scan(&Name)
	fmt.Print("Enter mobile: ")
	fmt.Scan(&Mobile)
	fmt.Print("Enter email: ")
	fmt.Scan(&Email)
	fmt.Print("Enter locality: ")
	fmt.Scan(&Locality)
	fmt.Print("Enter pin code: ")
	fmt.Scan(&Pin)
enterState:
	fmt.Print("Enter state: ")
	fmt.Scan(&State)
	State = strings.ToUpper(State)
	if validation.IsValidState(State) {
		goto enterCountry
	} else {
		fmt.Print("\nNot a valid state. Please try again.\n")
		goto enterState
	}
enterCountry:
	fmt.Print("Enter country: ")
	fmt.Scan(&Country)
	Country = strings.ToUpper(Country)
	if validation.IsValidCountry(Country) {
		goto validateStateCountry
	} else {
		fmt.Print("\nNot a valid Country. Please try again.\n")
		goto enterCountry
	}
validateStateCountry:
	if validation.IsValidStateForCountry(Country, State, validStates) {
		goto dataentry
	} else {
		fmt.Print("\nYour state is not in the country you entered. Please try again.\n")
		goto enterState
	}
dataentry:
	details := data.ID{
		Name:   Name,
		Mobile: Mobile,
		Email:  Email,
		CurrentAddress: data.Address{
			Locality: Locality,
			State:    State,
			Pin:      Pin,
			Country:  data.Country(Country),
		},
		PermanentAddress: data.Address{
			Locality: Locality,
			State:    State,
			Pin:      Pin,
			Country:  data.Country(Country),
		},
	}
	database.InsertDataDB(details)
	return details
}
