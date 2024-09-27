package validation

import (
	"strings"
)

func IsValidCountry(country string) bool {
	switch country {
	case "INDIA", "USA", "UK", "CANADA":
		return true
	default:
		return false
	}
}

func IsValidState(state string) bool {
	switch state {
	case "TELANGANA", "MAHARASHTRA", "KARNATAKA", "TEXAS", "TORONTO", "LONDON":
		return true
	default:
		return false
	}

}

func IsValidStateForCountry(country string, state string, validStates map[string][]string) bool {
	// Get the states for the given country
	country = strings.ToUpper(country)
	states, exists := validStates[country]
	if !exists {
		return false
	}

	// Check if the state exists in the list of valid states
	for _, s := range states {
		s = strings.ToUpper(s)
		if s == state {
			return true
		}
	}
	return false
}
