package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/iqubalreyaz/GolangLearning/structs/data"
	_ "github.com/mattn/go-sqlite3"
)

func CreateDB() {
	// Connect to SQLite database (creates the database if it doesn't exist)
	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create users table with all required fields
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		"id" INTEGER AUTOINCREMENT,		
		"name" TEXT,
		"mobile" INTEGER NOT NULL PRIMARY KEY,
		"email" TEXT UNIQUE,
		"permanent_locality" TEXT,
		"permanent_state" TEXT,
		"permanent_pin" INTEGER,
		"permanent_country" TEXT,
		"current_locality" TEXT,
		"current_state" TEXT,
		"current_pin" INTEGER,
		"current_country" TEXT
	);`

	// Execute the SQL statement to create the table
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Table created successfully.")

	// // Query and display all users
	// rows, err := db.Query("SELECT id, name, mobile, email, permanent_locality, permanent_state, permanent_pin, permanent_country, current_locality, current_state, current_pin, current_country FROM users")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer rows.Close()

	// // Iterate through the rows and display them
	// fmt.Println("Users in database:")
	// for rows.Next() {
	// 	var id int
	// 	var name string
	// 	var mobile int
	// 	var email string
	// 	var permanentLocality string
	// 	var permanentState string
	// 	var permanentPin int
	// 	var permanentCountry string
	// 	var currentLocality string
	// 	var currentState string
	// 	var currentPin int
	// 	var currentCountry string

	// 	err = rows.Scan(&id, &name, &mobile, &email, &permanentLocality, &permanentState, &permanentPin, &permanentCountry, &currentLocality, &currentState, &currentPin, &currentCountry)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Printf(
	// 		"ID: %d, Name: %s, Mobile: %d, Email: %s\nPermanent Address: %s, %s, %d, %s\nCurrent Address: %s, %s, %d, %s\n\n",
	// 		id, name, mobile, email, permanentLocality, permanentState, permanentPin, permanentCountry,
	// 		currentLocality, currentState, currentPin, currentCountry,
	// 	)
	// }

	// // Error handling for row iteration
	// err = rows.Err()
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

func InsertDataDB(details data.ID) {
	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// Insert a sample user into the table
	insertUserSQL := `
	INSERT INTO users(
		name, mobile, email,
		permanent_locality, permanent_state, permanent_pin, permanent_country,
		current_locality, current_state, current_pin, current_country
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	// Prepare the statement
	statement, err := db.Prepare(insertUserSQL)
	if err != nil {
		log.Fatal(err)
	}

	// Inserting user data
	_, err = statement.Exec(
		details.Name, details.Mobile, details.Email, details.PermanentAddress.Locality, details.PermanentAddress.State, details.PermanentAddress.Pin,
		details.PermanentAddress.Country, details.CurrentAddress.Locality, details.CurrentAddress.State, details.CurrentAddress.Pin, details.CurrentAddress.Country,
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User inserted successfully.")

}

func QueryUser() data.ID {
	var res data.ID
	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// Query and display all users
	rows, err := db.Query("SELECT id, name, mobile, email, permanent_locality, permanent_state, permanent_pin, permanent_country, current_locality, current_state, current_pin, current_country FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Iterate through the rows and display them
	fmt.Println("Users in database:")
	for rows.Next() {
		var id int
		var name string
		var mobile int
		var email string
		var permanentLocality string
		var permanentState string
		var permanentPin int
		var permanentCountry string
		var currentLocality string
		var currentState string
		var currentPin int
		var currentCountry string

		err = rows.Scan(&id, &name, &mobile, &email, &permanentLocality, &permanentState, &permanentPin, &permanentCountry, &currentLocality, &currentState, &currentPin, &currentCountry)
		if err != nil {
			log.Fatal(err)
		}

		res = data.ID{
			Name:   name,
			Mobile: mobile,
			Email:  email,
			PermanentAddress: data.Address{
				Locality: permanentLocality,
				State:    permanentState,
				Pin:      permanentPin,
				Country:  data.Country(permanentCountry),
			},
			CurrentAddress: data.Address{
				Locality: currentLocality,
				State:    currentState,
				Pin:      currentPin,
				Country:  data.Country(currentCountry),
			},
		}
	}

	// Error handling for row iteration
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return res
}
