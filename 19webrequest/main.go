package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)
 
const url = "https://github.com"
func main()  {
	fmt.Println("Lco web request")

	response,err := http.Get(url)

	if err!=nil{
		panic(err)
	}
	fmt.Printf("%T",response)

	defer response.Body.Close()

	databyte,err   :=  ioutil.ReadAll(response.Body)

	if err!=nil{
		panic(err)
	}
	content:=string(databyte)

	fmt.Println(content)
}