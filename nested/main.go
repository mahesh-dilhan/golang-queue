package main

import "fmt"

type Contact struct {
	phone string
}

type Person struct {
	name    string
	contact Contact
}

func (c *Contact) changeContact(newContact string) {
	c.phone = newContact
}
func main() {
	c := Contact{
		phone: "10098387123",
	}
	p := Person{
		name:    "Mahesh",
		contact: c,
	}

	fmt.Println(p)

	p.contact.changeContact("876123709123")

	fmt.Println(p)
}
