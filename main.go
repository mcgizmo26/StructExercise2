package main

import "fmt"

type contactInfo struct {
	email   string
	street  string
	city    string
	state   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Halpert",
		contact: contactInfo{
			email:   "jimHalpert@dundermiflin.com",
			street:  "1725 Slough Avenue",
			city:    "Scranton",
			state:   "PA",
			zipCode: 18505,
		},
	}

	// Create a new variable and point it to the original location of the jim struct.
	// address    value
	// jimPointer := &jim
	// jimPointer.updateName("Jimmy")
	// jim.print()

	// Use as a shortcut as long as in the function reciever you use *typeof ie. *person
	jim.updateName("Jimmy")
	jim.print()
}

// The reciever acts like a method
func (p person) print() {
	fmt.Printf("My name is %s %s and my email is %s and my business adress is %s %s, %s %v ", p.firstName, p.lastName, p.contact.email, p.contact.street, p.contact.city, p.contact.state, p.contact.zipCode)
}

// The *type (Type Descriptor) in the reciever means it type of a pointer to person.
func (pointerToPerson *person) updateName(newFirstName string) {
	// The *variable accesses the value of at that memory address.
	(*pointerToPerson).firstName = newFirstName
}

// Notes: Go is a pass by value language. You must use pointers to update the
// original struct.

// To turn an address into a value use *address.
// To turn a value into an address use &value.
