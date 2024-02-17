package main

import "fmt"

func main() {

	samar := User{"Samar","samar@gmail.com",21}
	fmt.Println(samar)

	fmt.Printf("Name of user is %v \n",samar.Name)
}

type User struct {
	Name  string
	email string
	age   int
}
