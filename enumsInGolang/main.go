package main

import "fmt"

// SIMPLE CONSTANTS IMPLEMENTATION
const (
	SUNDAY    = "Sunday"
	MONDAY    = "Monday"
	TUESDAY   = "Tuesday"
	WEDNESDAY = "Wednesday"
	THURSDAY  = "Thursday"
	FRIDAY    = "Friday"
	SATURDAY  = "Saturday"
)

//CONTANTS IMPLEMENTATION WITH IOTA
const (
	JANUARY  = iota // January is 0
	FEBRUARY        // February is 1 (increment 1 from last value)
	MARCH           // March is 2 (increment 1 from last value)
	APRIL
	MAY
	JUNE
	JULY
	AUGUST
	SEPTEMBER
	OCTOBER
	NOVEMBER
	DECEMBER
)

func main() {
	day := WEDNESDAY
	fmt.Println("Today is", day) // "Today is Monday"
	month := SEPTEMBER
	fmt.Println("\nMonth is", month+1) // Month is 9
}
