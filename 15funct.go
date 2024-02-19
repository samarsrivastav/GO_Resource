package main

import "fmt"
//we can place the function before or after main function 
func main()  {
	fmt.Println(" Functions ")
	first();
	fmt.Println(second(20))
	fmt.Println(proAdder(20,23,2344,4,32,211,1,2))

}
func first()  { //void funtion
	fmt.Println("1st function")
}
func second(a int) int { //int function can use anything 
	return a
}
func proAdder(values ...int) int {
	total :=0

	for _,val :=range values{
		total+=val
	}
	return total
}