package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

)
func main()  {
	fmt.Println("enter string")

	reader:=bufio.NewReader(os.Stdin)
	
	input,_ :=reader.ReadString('\n')

	numer,err:=strconv.ParseFloat(strings.TrimSpace(input),64)

	if err !=nil{
		fmt.Println(err)
	}else {
		fmt.Println("Your rating +1 is: ", numer+1)
	}
}