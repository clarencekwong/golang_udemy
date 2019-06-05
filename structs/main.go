package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Pop",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 90210,
		},
	}
	jimPointer := &jim
	jimPointer.updateName("jimmy")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}