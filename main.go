// first line of the program must be package main, it is the entry point of the program
package main

import (
	"fmt"
	"log"
)

// package level variable
var sGlobal = "John"

func main() { // main function is the entry point of the program, it is the first function to be executed, it is the only function that can be executed without being called

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
