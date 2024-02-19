package main

import "fmt"

func main()  {


	defer  fmt.Println("world2") //always print  it  in the lase of the funtion
	//defer follows last in first out
	fmt.Println("hello")

	defer  fmt.Println("world")
	defer  fmt.Println("world1")
}
//hello ->world1->world ->world2
