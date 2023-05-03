package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
	// or use just contactInfo
}

// var movsum person;

func main() {
	movsum := person{
		firstName: "Movsum",
		lastName:  "Aghakishiyev",
		contact: contactInfo{
			email:   "movsumaghakishiyev@gmail.com",
			zipCode: 1090,
		}}

	movsum.print()

	// movsumPointer := &movsum
	// movsumPointer.updateName("M & M")
	movsum.updateName("M & M")

	// fmt.Println("\n\n", movsumPointer, "\n")

	movsum.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p person) getFirstName() string {
	return p.firstName
}

func (pointerPerson *person) updateName(newFirstName string) {
	(*pointerPerson).firstName = newFirstName
}
