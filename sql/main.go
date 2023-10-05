package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func initDB() {
	// Connect to db.
	connStr := "user=user dbname=coffee_shop password=password2 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// Check if the connection is successful.
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Query.
	rows, err := db.Query("SELECT username FROM users;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Print out the rows.
	for rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(username)
	}

	// Close connection to db.
	defer db.Close()
}

func main() {
	initDB()
}
