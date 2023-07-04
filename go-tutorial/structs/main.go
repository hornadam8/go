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
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	jimptr := &jim
	fmt.Println(&jimptr)
	printPointer(jimptr)
	jim.updateName("Jimmy")
	jim.print()
}

func (personPtr *person) updateName(newFirstName string) {
	(*personPtr).firstName = newFirstName
}

func printPointer(ptr *person) {
	fmt.Println(&ptr)
}
func (personPtr *person) print() {
	fmt.Printf("%+v", *personPtr)
}
