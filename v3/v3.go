package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Person struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func main() {
	var contacts []Person
	readData(&contacts)
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
			addContact(&contacts, name, phone)
		case 2:
			fmt.Println(getAllContacts(contacts))
		case 3:
			fmt.Println("Enter contact name to remove ...")
			contactToRemove := readInput()
			removeContact(&contacts, contactToRemove)
		case 4:
			fmt.Println("Enter contact to update ...")
			contactToUpdate := readInput()
			fmt.Println("Enter new contact name ...")
			name := readInput()
			fmt.Println("Enter new contact phone ...")
			phone := readInput()
			updateContact(&contacts, contactToUpdate, name, phone)
		default:
			fmt.Println("Wrong case")
		}
	}
}

func addContact(contacts *[]Person, name, phone string) {
	newPerson := Person{Name: name, Phone: phone}
	*contacts = append(*contacts, newPerson)
	writeData(contacts)
}

func getIndexOfContact(contacts []Person, name string) int {
	for ind, val := range contacts {
		if val.Name == name {
			return ind
		}
	}
	return -1
}

func removeContact(contacts *[]Person, name string) bool {
	index := getIndexOfContact(*contacts, name)
	if index != -1 {
		*contacts = append((*contacts)[:index], (*contacts)[index+1:]...)
		writeData(contacts)
		return true
	}
	fmt.Println("Your contact does not match")
	return false
}

func updateContact(contacts *[]Person, name, updateName, updatePhone string) bool {
	index := getIndexOfContact(*contacts, name)
	if index != -1 {
		(*contacts)[index].Name = updateName
		(*contacts)[index].Phone = updatePhone
		writeData(contacts)
		return true
	}
	fmt.Println("Your contact does not match")
	return false
}

func getAllContacts(contacts []Person) []Person {
	return contacts
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

func readData(contacts *[]Person) {
	file, err := os.Open("v3/contact.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(contacts); err != nil {
		fmt.Println("Error while decoding file:", err)
		return
	}
}

func writeData(contacts *[]Person) {
	data, err := json.Marshal(contacts)
	if err != nil {
		fmt.Println("Error marshaling the data:", err)
		return
	}
	file, err := os.Create("v3/contact.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("Data written to file successfully.")
}
