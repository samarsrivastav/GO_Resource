package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("file handling")
	fmt.Print("\n\n")

	content := "this is the message"

	file, err := os.Create("./myfile.txt") //this will create a file

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content) //this will add the content in the file

	if err != nil {
		panic(err)
	}
	fmt.Println(length)
	defer file.Close()
	readFile("./myfile.txt")
}
// this  function reads the file content 
func readFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
