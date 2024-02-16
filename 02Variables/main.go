package main

import "fmt"

func main() {
	// 3 ways to declare a variable
	
	// m-1
	name := "samar"
	islogged := false
	floatex := 1.3214
	fmt.Println(name, " ", islogged, " ", floatex)

	// m-2 (i  will use this in most of the cases)
	var name1 = "samar"
	var islogged1 = false
	var floatex1 = 1.3214
	fmt.Println(name1, " ", islogged1, " ", floatex1)

	// m-3
	var name2 string="kjhsadf"
	fmt.Print(name2)
}
