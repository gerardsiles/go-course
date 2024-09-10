package user

import (
	"errors"
	"fmt"
	"time"
)

// Upper case makes it available, for type and for variables, like FirstName instead of firstName
type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

// Inheritance in structs
type Admin struct {
	email    string
	password string
	User
}

func NewAdmin(email, password, firstName, lastName, birthdate string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: firstName,
			lastName:  lastName,
			birthdate: birthdate,
			createdAt: time.Now(),
		},
	}
}

func (u User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func (u *User) ClearUserName() {
	u.birthdate = ""
	u.firstName = ""
	u.lastName = ""

}

func NewUser(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New(("first name, last name and birthdate are required"))
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}
