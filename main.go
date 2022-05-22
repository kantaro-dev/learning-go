package main

import "log"

func main() {
	// for i := 0; i <= 5; i++ {
	// 	log.Println(i)
	// }

	// animals := []string{"dog", "fish", "horse", "cow", "cat"}

	// for _, animal := range animals {
	// 	log.Println(animal)
	// }

	// animals := make(map[string]string)
	// animals["dog"] = "Fido"
	// animals["cat"] = "Fulffy"

	// for animalType, animal := range animals {
	// 	log.Println(animalType, animal)
	// }

	// var firstLine = "Once upon a midnight dereary"

	// for i, l := range firstLine {
	// 	log.Println(i, ":", l)
	// }

	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	var users []User

	users = append(users, User{"John", "Smith", "john@smith.com", 30})
	users = append(users, User{"Mary", "Jones", "mary@jones.com", 15})
	users = append(users, User{"Sally", "Browh", "sally@brown.com", 40})
	users = append(users, User{"Alex", "Smith", "alex@smith.com", 23})

	for _, l := range users {
		log.Println(l.FirstName, l.LastName, l.Email, l.Age)
	}

}
