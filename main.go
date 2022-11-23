// first line of the program must be package main, it is the entry point of the program
package main

import "fmt"

// package level variable
var name = "John"

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
	whatWasSaid, theOtherThing := saySomething()

	fmt.Println(whatWasSaid, theOtherThing)

}

func saySomething() (string, string) {
	return "Hello", "else"
}
