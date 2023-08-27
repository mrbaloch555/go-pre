package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

type Person struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func main() {
	var contacts []Person
	fmt.Println(contacts)

	for {
		fmt.Println("=================================================")
		fmt.Println("Choose your option...")
		fmt.Println("1 to Add new contact")
		fmt.Println("2 to Get all contacts")
		fmt.Println("3 to Remove a contact")
		fmt.Println("4 to Update a contact")
		fmt.Println("=================================================")
		fmt.Println("Enter your choice...")
		input := readInput()
		number, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Error converting to a number:", err)
		}

		switch number {
		case 1:
			fmt.Println("Enter new contact name ...")
			name := readInput()
			fmt.Println("Enter new contact phone ...")
			phone := readInput()
			addContact(dbConfig(), name, phone)
		case 2:
			docs, _ := readAllData(dbConfig())
			fmt.Println(docs)
		case 3:
			fmt.Println("Enter contact name to remove ...")
			contactToRemove := readInput()
			removeContact(dbConfig(), contactToRemove)
		case 4:
			fmt.Println("Enter contact to update ...")
			contactToUpdate := readInput()
			fmt.Println("Enter new contact name ...")
			name := readInput()
			fmt.Println("Enter new contact phone ...")
			phone := readInput()
			updateContact(dbConfig(), contactToUpdate, name, phone)
		default:
			fmt.Println("Wrong case")
		}
	}
}

func dbConfig() *sql.DB {
	db, err := sql.Open("sqlite3", "./contact.sql")
	if err != nil {
		return nil
	}

	tableExists := tableExists(db, "contacts")
	if !tableExists {
		createTableSQL := `
            CREATE TABLE contacts (
                id INTEGER PRIMARY KEY AUTOINCREMENT,
                name TEXT,
                phone TEXT
            )`
		_, err := db.Exec(createTableSQL)
		if err != nil {
			return nil
		}
	}

	return db
}

func tableExists(db *sql.DB, tableName string) bool {
	query := "SELECT name FROM sqlite_master WHERE type='table' AND name=?"
	var name string
	err := db.QueryRow(query, tableName).Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		return true
	}
	return name == tableName
}

func addContact(db *sql.DB, name, phone string) {
	defer db.Close()

	insertSQL := "INSERT INTO contacts (name, phone) VALUES (?, ?)"
	_, err := db.Exec(insertSQL, name, phone)
	if err != nil {
		fmt.Println("Error inserting data:", err)
		return
	}
}

func removeContact(db *sql.DB, name string) {
	defer db.Close()

	insertSQL := "DELETE FROM contacts where name=?"
	_, err := db.Exec(insertSQL, name)
	if err != nil {
		fmt.Println("Error deleting data:", err)
	}
}

func updateContact(db *sql.DB, name, updateName, updatePhone string) {
	defer db.Close()

	insertSQL := "UPDATE contacts SET name=?, phone=? where name=?"
	_, err := db.Exec(insertSQL, updateName, updatePhone, name)
	if err != nil {
		fmt.Println("Error updating data:", err)
	}
}

func readInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		return scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
	return ""
}

func readAllData(db *sql.DB) ([]Person, error) {
	defer db.Close()
	var contacts []Person

	rows, err := db.Query("SELECT id, name, phone FROM contacts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var person Person
		err := rows.Scan(&person.ID, &person.Name, &person.Phone)
		if err != nil {
			return nil, err
		}
		contacts = append(contacts, person)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return contacts, nil
}
