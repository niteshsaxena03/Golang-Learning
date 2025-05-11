package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}

func main() {
	fmt.Println("Started json in go")

	var person Person
	person.Name = "Nitesh"
	person.Age = 22
	person.IsAdult = true

	fmt.Println(person)

	//convert to json
	var jsonData, err = json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling json", err)
		return
	}
	fmt.Println(string(jsonData))

	//decoding json
	var personData Person
	err = json.Unmarshal(jsonData, &personData)
	if err != nil {
		fmt.Println("Error unmarshalling json", err)
		return
	}
	fmt.Println(personData)
}
