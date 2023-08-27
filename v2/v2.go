package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Contacts struct {
	Contacts []Person
}

type Person struct {
	Name  string
	Phone string
}

func main() {
	var conatct Contacts
	for {
		fmt.Println("=================================================")
		fmt.Println("Choose your option...")
		fmt.Println("1 to Add new contact")
		fmt.Println("2 to Get all contacts")
		fmt.Println("3 to Remove a contact")
		fmt.Println("4 to Update a contact")
		fmt.Println("=================================================")
		fmt.Println("Enter your choice...")
		input := ReadInput()
		number, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Error converting to a number:", err)
		}
		switch number {
		case 1:
			fmt.Println("Enter new contact name ...")
			name := ReadInput()
			fmt.Println("Enter new contact phone ...")
			phone := ReadInput()
			conatct.Add(name, phone)
			break
		case 2:
			fmt.Println(conatct.GetAll())
		case 3:
			fmt.Println("Enter contact name to remove ...")
			contactToRemove := ReadInput()
			conatct.Delete(contactToRemove)
		case 4:
			fmt.Println("Enter contact to update ...")
			conatctToUpdate := ReadInput()
			fmt.Println("Enter new contact name ...")
			name := ReadInput()
			fmt.Println("Enter new contact name ...")
			phone := ReadInput()
			conatct.Update(conatctToUpdate, name, phone)
		default:
			fmt.Println("Wrong case")
		}

	}
}

func (c *Contacts) Add(name string, phone string) {
	newPerson := Person{Name: name, Phone: phone}
	c.Contacts = append(c.Contacts, newPerson)
}

func (c *Contacts) GetIndex(name string) int {
	index := -1
	for ind, val := range c.Contacts {
		if val.Name == name {
			index = ind
		}
	}
	return index
}

func (c *Contacts) Delete(name string) bool {
	index := c.GetIndex(name)
	if index != -1 {
		c.Contacts = append(c.Contacts[:index], c.Contacts[index+1:]...)
		return true
	} else {
		fmt.Println("Your contact do not matched")
		return false
	}
}

func (c *Contacts) Update(name string, updateName string, updatePhone string) bool {
	index := c.GetIndex(name)
	if index != -1 {
		c.Contacts[index].Name = updateName
		c.Contacts[index].Phone = updatePhone
		return true
	} else {
		fmt.Println("Your contact do not matched")
		return false
	}
}

func (c *Contacts) GetAll() []Person {
	return c.Contacts
}

func ReadInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		return scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}

	return ""
}
