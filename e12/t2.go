package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type person struct {
	Firstname string  `json:"fname" xml:"fname"`
	Lastname  string  `json:"lname" xml:"lname"`
	Age       int     `json:"age" xml:"age"`
	Address   address `json:"addr" xml:"addr"`
}

type address struct {
	Street string `json:"street" xml:"street"`
	City   string `json:"city" xml:"city"`
}

func main() {
	p1 := person{"Vinod", "Kumar", 47, address{}}
	p2 := person{"Shyam", "Sundar", 48, address{}}

	p1.Address = address{"ISRO lyt", "Bangalore"}

	people := []person{p1, p2}

	out, _ := json.Marshal(people)
	fmt.Println(string(out))

	p1Json := `{"fname":"Vinod","lname":"Kumar","age":47}`
	var p3 = person{}
	json.Unmarshal([]byte(p1Json), &p3)
	fmt.Println(p3)

	xmlOut, _ := xml.MarshalIndent(p1, "", "   ")
	fmt.Println(string(xmlOut))
}
