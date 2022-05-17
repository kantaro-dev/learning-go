package main

import (
	"log"
	"time"
)

// var s = "seven"

// var firstName string
// var lastName string
// var phoneNumber string
// var age int
// var birthDate time.Time

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	// var s2 = "six"

	// log.Println("s is", s)
	// log.Println("s2 is", s2)

	// saySomething("xxx")

	user := User{
		FirstName:   "Blaise",
		LastName:    "Pascal",
		PhoneNumber: "123-456-7890",
	}

	log.Println(user.FirstName, user.LastName, "BirthDate:", user.BirthDate)
}

// func saySomething(s3 string) (string, string) {
// 	log.Println("s from the saySomething func is", s)
// 	return s3, "World"
// }
