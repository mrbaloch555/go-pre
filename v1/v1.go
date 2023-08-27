package main

import (
	"fmt"
	"strconv"
)

type Contacts struct {
	Contacts []string
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
			newConatc := ReadInput()
			conatct.Add(newConatc)
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
			fmt.Println("Enter new contact naem ...")
			newName := ReadInput()
			conatct.Update(conatctToUpdate, newName)
		default:
			fmt.Println("Wrong case")
		}

	}
}

func (c *Contacts) Add(name string) bool {
	c.Contacts = append(c.Contacts, name)
	return true
}

func (c *Contacts) GetIndex(name string) int {
	index := -1
	for ind, val := range c.Contacts {
		if val == name {
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

func (c *Contacts) Update(name string, update string) bool {
	index := c.GetIndex(name)
	if index != -1 {
		c.Contacts[index] = update
		return true
	} else {
		fmt.Println("Your contact do not matched")
		return false
	}
}

func (c *Contacts) GetAll() []string {
	return c.Contacts
}

func ReadInput() string {
	var userInput string
	_, err := fmt.Scan(&userInput)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return ""
	}
	return userInput
}
