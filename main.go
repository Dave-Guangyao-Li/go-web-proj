// first line of the program must be package main, it is the entry point of the program
package main

import (
	"fmt"
	"log"
	"time"
)

// package level variable
var sGlobal = "John"

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

	// declare a variable
	var whatToSay string
	var i int // int type depends on the system

	whatToSay = "Goodbye"

	// fmt stand for format, it is a package
	fmt.Println(whatToSay)

	i = 10

	fmt.Println("i is", i)

	// := is a short declaration operator, it is used to declare and initialize a variable
	// whatWasSaid, theOtherThing := saySomething()

	// fmt.Println(whatWasSaid, theOtherThing)

	// pointer is a variable that stores the memory address of another variable
	var myString string = "Blue"

	log.Println("myString is set to", myString)

	// & for reference
	changeUsingPointer(&myString) // & means get the address of the variable
	log.Println("myString is set to", myString)

	var sLocal = "seven"
	log.Println("sGlobal is", sGlobal) // global variable
	log.Println("sLocal is", sLocal)   // local variable
	saySomething("xxx")

	// var firstName string
	// var lastName string
	// var age int
	// var phoneNumber string
	// var birthDate time.Time

	// lowercase variable is not exported, capital letter variable is exported
	user := User{
		FirstName:   "John",
		LastName:    "Doe",
		Age:         30,
		PhoneNumber: "123456789",
		BirthDate:   time.Now(),
	}

	log.Println(user.FirstName, user.LastName, user.Age, user.PhoneNumber, user.BirthDate)

}

func saySomething(s3 string) (string, string) {
	log.Println("s from the func is", sGlobal)
	return s3, "else"
}

func changeUsingPointer(s *string) {
	log.Println("s is set to", s)
	newValue := "Red"
	*s = newValue // chane the value of the variable that the pointer is pointing to
}
