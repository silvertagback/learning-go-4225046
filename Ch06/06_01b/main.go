package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Files")
	fileName := "./fromString.txt"
	file, err := os.Create(fileName)
	defer file.Close()
	checkError(err)
	length, err := io.WriteString(file, "Hello from Go!")
	fmt.Printf("File Character Length: %d\n", length)
	readFile(fileName)
}

func readFile(fileName string) {
	data, err := os.ReadFile(fileName)
	checkError(err)
	fmt.Println("File Contents:", string(data))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}