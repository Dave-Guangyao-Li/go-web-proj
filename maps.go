package main

import (
	"log"
	"sort"
)

type User struct {
	FirstName string
	LastName  string
}

func main() {
	// slice is a dynamic array, it is a reference type, it is a composite type
	var mySlice []string
	numbers := []int{3, 2, 1, 4, 5}
	names := []string{"John", "Mary", "Bob", "Alice"}
	mySlice = append(mySlice, "Alice")
	mySlice = append(mySlice, "Charlie")
	mySlice = append(mySlice, "Bob")
	log.Println(mySlice)
	sort.Strings(mySlice)
	log.Println("sorted: ", mySlice)

	log.Println("numbers: ", numbers)
	sort.Ints(numbers)
	log.Println("sorted numbers numbers[0:2]: ", numbers[0:2])

	log.Println("names: ", names)
	// use short declaration operator to declare and initialize a variable
	// myMap := make(map[string]string)
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
