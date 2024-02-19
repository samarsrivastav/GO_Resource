package main

import "fmt"

func main() {

	samar := User{"Samar","samar@gmail.com",21}
	fmt.Println(samar)

	fmt.Printf("Name of user is %v \n",samar.Name)
	
	samar.GetStatus()
	samar.NewMail()
	fmt.Printf("Name of user is %v \n",samar.email)
}

type User struct {
	Name  string
	email string
	age   int
}
 //method getter
func (u User) GetStatus() {
	fmt.Println("email of user: ",u.email)
}
func (u User) NewMail() {
	u.email="sdfa@adsfk.com"
	fmt.Println("email of user: ",u.email) //this wont change the original...but a copy....(call by value)
}