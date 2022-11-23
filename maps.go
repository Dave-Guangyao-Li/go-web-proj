package main

import (
	"log"
)

type User struct {
	FirstName string
	LastName  string
}

func main() {

	// use short declaration operator to declare and initialize a variable
	// myMap := make(map[string]string)
	myMap1 := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	log.Println(myMap1)

	myMap := make(map[string]int)
	// myMap["dog"] = "bark"
	// myMap["cat"] = "meow"

	// myMap["dog"] = "woof"
	myMap["first"] = 1
	// log.Println(myMap["dog"])
	// log.Println(myMap["cat"])
	log.Println(myMap["first"])

	me := User{
		FirstName: "John",
		LastName:  "Doe",
	}

	userMap := make(map[string]User)
	// map is unordered, it is not indexed, it is a collection of key value pairs

	userMap["me"] = me

	log.Println(userMap["me"].FirstName)
}
