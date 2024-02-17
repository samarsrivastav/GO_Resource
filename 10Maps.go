package main

import "fmt"

func main() {
	language := make(map[string]string)

	language["js"] = "js"
	language["py"] = "py"
	language["java"] = "java"

	fmt.Println(language)
	fmt.Println(language["py"])

}
