// first line of the program must be package main, it is the entry point of the program
package main

import (
	"log"
	"time"
)

// struct is a collection of fields, it is a user defined type, it is a composite type
type User struct {
	FirstName   string
	LastName    string
	Age         int
	PhoneNumber string
	BirthDate   time.Time
}

type myStruct struct {
	FirstName string
}

// receiver function, it is a function that is attached to a type
func (m *myStruct) printFirstName() string {
	return m.FirstName
}

func main() { // main function is the entry point of the program, it is the first function to be executed, it is the only function that can be executed without being called

	// call receiver function on a struct
	var myVar myStruct
	myVar.FirstName = "John"
	myVar2 := myStruct{FirstName: "Mary"}
	log.Println(myVar2.printFirstName())
	log.Println("myVar is set to", myVar.printFirstName())

	user := User{
		FirstName:   "John",
		LastName:    "Doe",
		Age:         30,
		PhoneNumber: "123456789",
		BirthDate:   time.Now(),
	}

	log.Println(user.FirstName, user.LastName, user.Age, user.PhoneNumber, user.BirthDate)

}

func changeUsingPointer(s *string) {
	log.Println("s is set to", s)
	newValue := "Red"
	*s = newValue // chane the value of the variable that the pointer is pointing to
}
