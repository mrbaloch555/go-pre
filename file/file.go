package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("file/file.txt")
	if err != nil {
		fmt.Println("Error opening the file: ", err)
		return
	}

	defer file.Close()

	data := make([]byte, 1024)

	_, err = file.Read(data)
	if err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}
	fmt.Println(string(data))

	file, err = os.OpenFile("file/file.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println("Error opening the file: ", err)
		return
	}
	defer file.Close()

	newData := []byte("\nI am a software engineer")
	_, err = file.Write(newData)
	if err != nil {
		fmt.Println("Error writing to file: ", err)
		return
	}

	fmt.Println("Data appended to the file successfully")
	ReadAndWrite()
}

func ReadAndWrite() {

	file, err := os.OpenFile("file/file.txt", os.O_CREATE|os.O_RDWR, 0644)

	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}

	defer file.Close()

	data := make([]byte, 1024)

	n, err := file.Read(data)
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}
	fmt.Println("Read data: ", string(data[:n]))

	writeData := []byte(string(data) + "\nI live in lahore\n")
	// copy(data[n:], writeData)
	file.Write(writeData)

	// _, err = file.WriteAt(data, 0)

	if err != nil {
		fmt.Println("Error writing to file")
		return
	}

	fmt.Println("Data written back to the file successfully")

}
