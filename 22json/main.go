package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"` //this will show the key name as coursename
	Price    int
	Platform string `json:"website"` //this will show the key name as website
	Password string `json:"-"` //this will hide the password at all cases
	Tags    []string `json:"tags,omitempty"` //this will not show if nil
}

func main() {
	println(" JSON ")
	EncodeJson()
	DecodeJson()
}
func EncodeJson()  {
	lcoCourse := []course{
		{"a",12,"asdf","a12",[]string{"a1","b2"}},
		{"b",112,"fvvf","f123",[]string{"f1","g2"}},
		{"cx",112,"ewqr","g123",nil},
	}	

	// package this data as json data 

	finalJson,err := json.MarshalIndent(lcoCourse,"","\t")

    if err!=nil{
		panic(err)
	}
	println(string(finalJson))
}
func DecodeJson()  {
	jsondataFromweb := []byte(`
		{
			"coursename": "b",
			"Price": 112,
			"website": "fvvf",
			"tags": ["f1","g2"]
		}
	`)
	var lcoCourse course

	check := json.Valid(jsondataFromweb)
    if check{
		fmt.Printf("valid\n")
		json.Unmarshal(jsondataFromweb,&lcoCourse)
		fmt.Println(lcoCourse)

	}else{
		fmt.Printf("not valid json \n")
	}


	//map
	var myonlinedata map[string]interface{}
	json.Unmarshal(jsondataFromweb,&myonlinedata)
	// fmt.Println(myonlinedata) //or print using loop

	for key,value :=range myonlinedata{
		fmt.Printf("%v -> %v \n",key,value)
	}

}
