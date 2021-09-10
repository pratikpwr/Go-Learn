package main

import "time"

// Struct in Go
// by this way struct can be defined
type User struct {
	firstname string
	lastname  string
	birthdate string
	createdAt time.Time
}

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
