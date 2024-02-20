package main

import (
	"fmt"
	"net/url"
)

const myurl string ="https://samar.com:2000/id?name=samar&last=srivastav"
func main()  {
	fmt.Println(myurl) //print whole url
	result,err := url.Parse(myurl)

	if err!=nil{
		panic(err)
	}
	fmt.Println(result) //print the url
	fmt.Println(result.Scheme) //https
	fmt.Println(result.Path) //id
	fmt.Println(result.Host) //samar.com:2000
	fmt.Println(result.Port()) //2000
	
	fmt.Println(result.RawQuery)  //name=samar & last -srivasarc


	// to  extract  values 
	qparams := result.Query()
	fmt.Println(qparams["name"]) //[samar]

	for _,val := range qparams{
		fmt.Print(val)
	}


}