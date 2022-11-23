package main

import (
	"log"
	"sort"
)

func main() {
	// slice is a dynamic array, it is a reference type, it is a composite type
	mySlice1 := make([]string, 0, 5) // make a slice of strings, length is 0, capacity is 5
	mySlice1 = append(mySlice1, "Red", "Blue", "Green", "Yellow", "Pink", "Purple")
	log.Println(mySlice1)
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

}
