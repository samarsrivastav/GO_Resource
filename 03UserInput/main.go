package main

import (
	"bufio"
	"fmt"
	"os"
)
func main()  {
	fmt.Println("enter string")

	reader:= bufio.NewReader(os.Stdin)

	print("enter: ")
	input , _:=reader.ReadString('\n')
	println(input)
}