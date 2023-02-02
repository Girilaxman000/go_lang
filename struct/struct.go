package main

import (
	"fmt"
	"strconv"
)

// Define person struict
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// Greeting Method (value receiver)
// p is similar to this in js
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + "my birth date" + strconv.Itoa(p.age)
}

// hasBirthDay method (pointer receiver)
// changing something and doesn't have return type
func (p *Person) hasBirthDay() {
	p.age++
}

func main() {
	//same as class
	//different methods associated with that struct
	//two different types of methods: pointer receivers and value receivers

	//initialize a Person using struct
	person1 := Person{firstName: "Laxman", lastName: "giri", city: "rolpa", gender: "m", age: 24}
	// fmt.Println(person1)

	//value receivers calculating receiver
	//we didn't change anything just receive value and show it.
	person1.hasBirthDay()
	fmt.Println(person1.greet())

	//pointer receiver changing

}
