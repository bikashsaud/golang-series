package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("File Handling with go... ")
	fileContent := "Go language for serverless applications"
	filePathName := "./golang_book.txt"
	writeFile(filePathName, fileContent)
	readFile(filePathName)
}

func writeFile(fileName string, content string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("error while creating file: ", err)
	}
	len, err := io.WriteString(file, content)
	checkNilError(err)
	fmt.Println("length is : ", len)
	defer file.Close()
}

func readFile(filenamepath string) {
	fmt.Println("Read file.")
	databytes, err := ioutil.ReadFile(filenamepath)
	checkNilError(err)
	fmt.Println("File COntent:", databytes)
	fmt.Println("File COntent:", string(databytes))

}

// common way to handle error
func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
