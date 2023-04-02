package queries

import (
	"fmt"
	"testing"

	"github.com/Ba1vo/Proektirovanie/decoder"
)

type testObject struct {
	user   decoder.RegUser
	create bool
}

func TestEmptyUser(t *testing.T) {
	var testUsers = []testObject{
		{
			user: decoder.RegUser{
				Name:  "Joka",
				Email: "2012egor@gmail.com",
				Pass:  "7b1963714055e4cfd07bff3f91a48e085e6a02bf734ea40e24c86f8488e80333",
			},
			create: true,
		},
		{
			user: decoder.RegUser{
				Name:  "Boka",
				Email: "2055egor@gmail.com",
				Pass:  "7b1963714055e4cfd07bff3f91a48e085e6a02bf734ea40e24c86f8488e80333",
			},
			create: false,
		},
	}
	for _, user := range testUsers {
		if user.create {
			AddUser(user.user)
		}
		newUser, err := GetUser(decoder.AuthUser{Email: user.user.Email, Pass: user.user.Pass})
		if err != nil && err.Error() == "Empty" {
			fmt.Printf("User(Email: %s, Pass: %s) doesn't exist\n", user.user.Email, user.user.Pass)

		} else {
			fmt.Printf("User(Email: %s, Pass: %s) found. Full creds: ID: %d, Name: %s Email: %s, Pass: %s, Date: %s\n",
				user.user.Email, user.user.Pass, newUser.ID, newUser.Name, newUser.Email, newUser.Pass, newUser.Date)
		}
		if user.create {
			DeleteUser(newUser.ID)
		}
	}
}
