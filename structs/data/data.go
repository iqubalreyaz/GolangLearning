package data

type Country string
type State string

const (
	INDIA  Country = "India"
	USA    Country = "USA"
	UK     Country = "UK"
	CANADA Country = "Canada"

	TELANGANA   State = "Telangana"
	MAHARASHTRA State = "Maharashtra"
	KARNATAKA   State = "Karnataka"
	TEXAS       State = "Texas"
	LONDON      State = "London"
	TORONTO     State = "Toronto"
)

type ID struct {
	Name             string
	Mobile           int
	Email            string
	PermanentAddress Address
	CurrentAddress   Address
}

type Address struct {
	Locality string
	State    string
	Pin      int
	Country  Country
}

var (
	name     string
	mobile   int
	email    string
	locality string
	state    string
	pin      int
	country  string
)
