package main

import "fmt"

func main()  {
	fmt.Println("-- Pointers --") 

	number :=25

	ptr := &number

	fmt.Println("pointer value " ,ptr)
	fmt.Println("pointer value " ,*ptr)
	
	*ptr  +=2
	fmt.Println("pointer value " ,*ptr)

}