package main

import (
	"fmt"
	"strconv"
)

// Person .
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int

	// firstName, lastName, city, gender string
	// age                               int
}

func (p Person) greet() string {
	return p.firstName + " " + p.lastName + " and age: " + strconv.Itoa(p.age)
}

func (p *Person) hasBirthday() {
	p.age++
}

func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "M" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	fmt.Println("Struct")
	person1 := Person{
		firstName: "Wang",
		lastName:  "Fa",
		city:      "Saru",
		gender:    "M",
		age:       30,
	}

	person2 := Person{"Bob", "John", "New", "F", 30}
	fmt.Println(person1)

	person1.hasBirthday()
	person1.getMarried("Williams")
	person2.getMarried("Thompson")
	fmt.Println(person2.greet())
}
