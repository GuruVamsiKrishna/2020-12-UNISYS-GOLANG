package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type Address struct {
	Street string `json:"streetAddress" xml:"street"`
	City   string `json:"city" xml:"city"`
	State  string `json:"stateOrRegion" xml:"state"`
}

type Person struct {
	Name    string  `json:"name" xml:"name"`
	Email   string  `json:"emailAddress" xml:"email"`
	Address Address `json:"address" xml:"addr"`
}

func main() {
	p1 := Person{"Vinod", "vinod@vinod.co", Address{"1st cross", "Bangalore", "Karnataka"}}

	bytes, err := json.MarshalIndent(p1, "", "\t")
	fmt.Println(string(bytes))

	bytes, err = xml.MarshalIndent(p1, "", "   ")
	if err == nil {
		fmt.Println(string(bytes))
	}

	jsonStr := `{
		"name": "Vinod",
		"emailAddress": "vinod@vinod.co",
		"address": {
			"streetAddress": "1st cross",
			"city": "Bangalore",
			"stateOrRegion": "Karnataka"
		}
	}`

	var p2 Person
	json.Unmarshal([]byte(jsonStr), &p2)
	fmt.Println("p2 is", p2)

}
