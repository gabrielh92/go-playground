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
	gabe := person{firstName: "Gabriel",
		lastName: "Hruskovec",
		contactInfo: contactInfo{
			email:   "gabe@gabe.com",
			zipCode: 11206,
		},
	}
	gabe.updateName("gabe")
	gabe.print()
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
