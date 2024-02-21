package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("request")

	// get request
	// PerformGetRequest()

	//post request
	PerformPostRequest()
}

func PerformGetRequest() {
	const myurl = "https://jsonplaceholder.typicode.com/posts"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Print("Status code: ", response.Status)
	fmt.Print("\nStatus code: ", response.ContentLength, "\n")
	content, err := ioutil.ReadAll(response.Body)

	// method-1 to show the content (will use this one )
	// fmt.Println(string(content))

	// m-2
	var responseString strings.Builder
	bytecount, _ := responseString.Write(content)

	fmt.Print(bytecount)
	fmt.Print(responseString.String())
}

func PerformPostRequest() {
	const myurl = "https://jsonplaceholder.typicode.com/posts"

	requestBody := strings.NewReader(`
		{
			"course":"golang",
			"id":120
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))

}
