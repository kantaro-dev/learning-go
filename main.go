package main

import "log"

// import "sort"

// type User struct {
// 	FirstName string
// 	LastName  string
// }

func main() {
	// myMap := make(map[string]string)

	// myMap["dog"] = "Samson"
	// myMap["other-dog"] = "Cassie"
	// myMap["dog"] = "Fido"

	// log.Println(myMap["dog"])
	// log.Println(myMap["other-dog"])

	// myMap := make(map[string]int)

	// myMap["First"] = 1
	// myMap["Second"] = 2

	// log.Println(myMap["First"], myMap["Second"])

	// myMap := make(map[string]User)

	// me := User{
	// 	FirstName: "Rene",
	// 	LastName:  "Descartes",
	// }

	// myMap["me"] = me
	// log.Println(myMap["me"].FirstName)

	// var mySlice []string

	// mySlice = append(mySlice, "Rene")
	// mySlice = append(mySlice, "George")
	// mySlice = append(mySlice, "Alfred")

	// log.Println(mySlice)

	// var mySliceInt []int

	// mySliceInt = append(mySliceInt, 2)
	// mySliceInt = append(mySliceInt, 1)
	// mySliceInt = append(mySliceInt, 3)

	// sort.Ints(mySliceInt)

	// log.Println(mySliceInt)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	log.Println(numbers)

	log.Println(numbers[0:2])
	log.Println(numbers[6:9])

	names := []string{"one", "seven", "fish", "cat"}

	log.Println(names)
}
