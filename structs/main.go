package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

//2 types of methods: Value receivers, Pointer receivers

//value receiver
func (p Person) greet() string {
	return "Hello my name is " + p.firstName
}

//pointer reciever
func (p *Person) hasBirthday() bool {
	p.age++
	//will change the actual value of age for that user
}

func (p *Person) getMarried(spouseLastName string) {
	if p.gender == 'm' {
		return
	}
	else {
		p.lastName = spuspouseLastName
	}
}

func main() {
	person1 := Person{
		firstName: "ahmad",
		lastName:  "Smith",
		city:      "boston",
		gender:    "f",
		age:       25,
	}

	/*
		person1 := Person{
			"ahmad",
			"Smith",
			"boston",
			"M",
			25,
		}
	*/

	fmt.Println(person1.greet())
	fmt.Println(person1.age)
	person1.hasBirthday()
	fmt.Println(person1.age)
	person1.getMarried("williams")


	//greet is a function that accepts structs of
	// of type person
}
