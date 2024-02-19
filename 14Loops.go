package main

import "fmt"

func main() {
	fmt.Println("Loops ")

	days := []string{"mon", "tue", "wed"}

	fmt.Println(days)

	// loop type -1
	// for d:=0;d<len(days);d++{
	// 	fmt.Println(days[d])
	// }

	// range based loop
	for i := range days {
		fmt.Println(i) //it will output 0 1 2
	}

	// range based direct value

	for _, day := range days {
		fmt.Println(day) //this will output the days
	}

	// while loop syntax
	i :=0
	for i < 10 {
		fmt.Print(i," ")
		i++
	}
}
