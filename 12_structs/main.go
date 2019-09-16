package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// Greeting method (value receiver)
func (p Person) greet() string {
	return "Hello, I am " + p.firstName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "M" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// init person using struct
	person1 := Person{firstName: "Chris", lastName: "George", city: "St. Louis", gender: "M", age: 31}
	// person1 := Person{"Chris", "George", "St. Louis", "M", 31}
	fmt.Println(person1)
	fmt.Println(person1.firstName)
	person1.age++
	fmt.Println(person1)
	person1.hasBirthday()
	fmt.Println(person1.greet())
	person2 := Person{firstName: "Biscuits", lastName: "TheDog", city: "St. Louis", gender: "F", age: 22}
	fmt.Println(person2)
	person2.getMarried(person1.lastName)
	fmt.Println(person2)
}
