package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
}

type Contact struct {
	Email string
	Phone string
}
type Address struct {
	House string
	Area  string
	State string
}

type Employee struct {
	person  Person
	contact Contact
	address Address
}

func main() {
	var person1 Person
	person1.firstName = "Nitesh"
	person1.lastName = "Saxena"
	person1.age = 22

	fmt.Println(person1)

	var employee1 Employee
	employee1.person = person1
	employee1.contact.Email = "nitesh@gmail.com"
	employee1.contact.Phone = "1234567890"
	employee1.address.House = "57"
	employee1.address.Area = "Delhi"
	employee1.address.State = "Delhi"

	fmt.Println(employee1)
}
