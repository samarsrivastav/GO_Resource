package main

import (
	"fmt"
	"time"
)
 
func main()  {
	fmt.Println("Time")
	currtime :=time.Now()

	fmt.Println(currtime)

	// formatting 
	fmt.Println(currtime.Format("01-02-2006 Monday 15:04:05"))
}