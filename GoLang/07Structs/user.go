package main

import (
	"fmt"
	"time"
)

// Struct in Go
// by this way struct can be defined
type User struct {
	firstname string
	lastname  string
	birthdate string
	createdAt time.Time
}

// It will create a new copy of variable
// which will increase resourse consuption
func CreateUser(firstName string, lastName string, birthDate string) (user User) {
	created := time.Now()

	user = User{
		firstname: firstName,
		lastname:  lastName,
		birthdate: birthDate,
		createdAt: created,
	}
	return
}

// using the pointers to create User which will save space
// it will use same copy of  variable
func CreateUserUsingPointer(firstName string, lastName string, birthDate string) *User {
	created := time.Now()

	user := User{
		firstname: firstName,
		lastname:  lastName,
		birthdate: birthDate,
		createdAt: created,
	}
	return &user
}

// to connect a function to struct and convert it into method
// we use receiver argument as struct
// we can use these methods on objects like newUser.OutputData()

func (user User) outputData() {
	fmt.Printf("My name is %v %v, I am born on %v.", user.firstname, user.lastname, user.birthdate)
}

func outputUerData(user *User) {
	// go automatically dereferences pointer structs
	// else we had to do like (*user).firstname
	fmt.Printf("My name is %v %v, I am born on %v.", user.firstname, user.lastname, user.birthdate)
}
